package main

import (
	"log"
	"os"

	// Domain
	"apirest/internal/domain/user"

	// Application - Commands
	masterCmd "apirest/internal/application/master/command"
	userCmd "apirest/internal/application/user/command"

	// Application - Queries
	masterQuery "apirest/internal/application/master/query"
	userQuery "apirest/internal/application/user/query"

	// Infrastructure
	"apirest/internal/infrastructure/mail"
	"apirest/internal/infrastructure/messaging/rabbitmq"
	"apirest/internal/infrastructure/persistence/postgres"
	"apirest/internal/infrastructure/security/jwt"

	// Presentation
	"apirest/internal/presentation/http"
	"apirest/internal/presentation/http/middleware"
	checkCtrl "apirest/internal/presentation/http/rest/check"
	masterCtrl "apirest/internal/presentation/http/rest/master"
	userCtrl "apirest/internal/presentation/http/rest/user"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	log.Println("Starting Work Time Controller API...")

	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	//
	// INFRASTRUCTURE SETUP
	//

	// Database
	dbConfig := postgres.LoadConfigFromEnv()
	db, err := postgres.NewDatabase(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// JWT
	jwtSecret := getEnv("JWT_SECRET", "your-secret-key-change-in-production")
	jwtIssuer := getEnv("JWT_ISSUER", "worktic-api")
	jwtGenerator := jwt.NewGenerator(jwtSecret, jwtIssuer)
	jwtValidator := jwt.NewValidator(jwtSecret)

	// RabbitMQ (optional)
	var rabbitMQConn *rabbitmq.Connection
	rabbitMQURL := getEnv("RABBITMQ_URL", "")
	if rabbitMQURL != "" {
		rabbitMQConn, err = rabbitmq.NewConnection(rabbitmq.Config{URL: rabbitMQURL})
		if err != nil {
			log.Printf("Failed to connect to RabbitMQ: %v", err)
		} else {
			defer rabbitMQConn.Close()
		}
	}

	// Mail service
	mailConfig := mail.Config{
		Host:     getEnv("MAIL_HOST", "localhost"),
		Port:     1025,
		Username: getEnv("MAIL_USERNAME", ""),
		Password: getEnv("MAIL_PASSWORD", ""),
		From:     getEnv("MAIL_FROM", "noreply@worktic.com"),
	}
	mailer := mail.NewMailer(mailConfig)
	_ = mailer // Use it later

	//
	// REPOSITORIES
	//

	userRepo := postgres.NewPostgresUserRepository(db)

	// Master repositories
	masterRepo := postgres.NewPostgresMasterRepository(db)
	masterProfileRepo := postgres.NewPostgresMasterProfileRepository(db)
	masterAccessLogRepo := postgres.NewPostgresMasterAccessLogRepository(db)

	//
	// DOMAIN SERVICES
	//

	userService := user.NewService(userRepo)

	//
	// APPLICATION HANDLERS - COMMANDS
	//

	// User commands
	refreshTokenHandler := userCmd.NewRefreshTokenHandler(jwtValidator, jwtGenerator)
	logoutHandler := userCmd.NewLogoutHandler(masterAccessLogRepo)

	// Master commands
	createMasterHandler := masterCmd.NewCreateMasterHandler(userService, masterRepo, masterProfileRepo)
	loginMasterHandler := masterCmd.NewLoginMasterHandler(userService, masterRepo, masterAccessLogRepo, jwtGenerator)
	updateMasterProfileHandler := masterCmd.NewUpdateMasterProfileHandler(masterProfileRepo)

	//
	// APPLICATION HANDLERS - QUERIES
	//

	// User queries
	whoAmIHandler := userQuery.NewWhoAmIHandler(userRepo)

	// Master queries
	getMasterProfileHandler := masterQuery.NewGetMasterProfileHandler(masterRepo, masterProfileRepo, userRepo)

	//
	// CONTROLLERS
	//

	checkController := checkCtrl.NewCheckController(db)

	userAuthController := userCtrl.NewAuthController(
		refreshTokenHandler,
		logoutHandler,
		whoAmIHandler,
	)

	masterAuthController := masterCtrl.NewAuthController(loginMasterHandler)
	masterAccountController := masterCtrl.NewAccountController(
		getMasterProfileHandler,
		updateMasterProfileHandler,
	)
	masterUsersController := masterCtrl.NewUsersController(createMasterHandler)

	controllers := &http.Controllers{
		CheckController:         checkController,
		UserAuthController:      userAuthController,
		MasterAuthController:    masterAuthController,
		MasterAccountController: masterAccountController,
		MasterUsersController:   masterUsersController,
	}

	//
	// FIBER APP SETUP
	//

	app := fiber.New(fiber.Config{
		AppName:      "Work Time Controller API",
		ErrorHandler: customErrorHandler,
	})

	// Middleware
	authMiddleware := middleware.NewAuthMiddleware(jwtValidator)
	app.Use(middleware.Logger())
	app.Use(middleware.CORS())

	// Health check
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to Worktic service!",
			"status": true,
		})
	})

	// Health check
	app.Get("/api", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to Worktic API. Version 1 available at /api/v1",
			"status": true,
		})
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  true,
			"service": "worktic-api",
		})
	})

	// Setup routes
	http.SetupRoutes(app, controllers, authMiddleware)

	//
	// START SERVER
	//

	port := getEnv("PORT", "8080")
	log.Printf("Server starting on port %s", port)
	log.Printf("API available at: http://localhost:%s/api/v1", port)
	log.Printf("Health check at: http://localhost:%s/health", port)

	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func customErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	return c.Status(code).JSON(fiber.Map{
		"success": false,
		"error":   err.Error(),
	})
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
