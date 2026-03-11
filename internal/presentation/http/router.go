package http

import (
	"apirest/internal/presentation/http/middleware"
	checkCtrl "apirest/internal/presentation/http/rest/check"
	masterCtrl "apirest/internal/presentation/http/rest/master"
	userCtrl "apirest/internal/presentation/http/rest/user"

	"github.com/gofiber/fiber/v2"
)

type Controllers struct {
	// Check
	CheckController *checkCtrl.CheckController

	// User
	UserAuthController *userCtrl.AuthController

	// Master
	MasterAuthController    *masterCtrl.AuthController
	MasterAccountController *masterCtrl.AccountController
	MasterUsersController   *masterCtrl.UsersController
}

func SetupRoutes(app *fiber.App, controllers *Controllers, authMiddleware *middleware.AuthMiddleware) {
	// API v1
	api := app.Group("/api/v1")

	//
	// PUBLIC ROUTES
	//

	// Check endpoints
	api.Get("/check", controllers.CheckController.Check)
	api.Get("/check/database", controllers.CheckController.CheckDatabase)
	api.Get("/check/broker", controllers.CheckController.CheckBroker)
	api.Get("/check/mailer", controllers.CheckController.CheckMailer)
	api.Get("/check/all", controllers.CheckController.CheckAll)

	// Master login (public)
	api.Post("/master/auth/login", controllers.MasterAuthController.Login)

	//
	// AUTHENTICATED ROUTES (ANY ROLE)
	//

	authRoutes := api.Group("", authMiddleware.Authenticate())

	// Common auth endpoints
	authRoutes.Post("/auth/refresh", controllers.UserAuthController.RefreshToken)
	authRoutes.Post("/auth/logout", controllers.UserAuthController.Logout)
	authRoutes.Get("/auth/whoami", controllers.UserAuthController.WhoAmI)

	//
	// MASTER ROUTES
	//

	masterRoutes := api.Group("/master", authMiddleware.Authenticate(), authMiddleware.RequireRole("master"))

	// Master account
	masterRoutes.Get("/account/profile", controllers.MasterAccountController.GetProfile)
	masterRoutes.Patch("/account/profile", controllers.MasterAccountController.UpdateProfile)

	// Master user management
	masterRoutes.Post("/users", controllers.MasterUsersController.CreateMaster)
	masterRoutes.Get("/users", controllers.MasterUsersController.GetAllMasters)
	masterRoutes.Get("/users/:id/profile", controllers.MasterUsersController.GetMasterProfile)
	masterRoutes.Delete("/users/:id", controllers.MasterUsersController.DeleteMaster)

	//
	// EMPLOYEE ROUTES (TODO)
	//
}
