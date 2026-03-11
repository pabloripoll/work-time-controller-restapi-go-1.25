package check

import (
	"apirest/internal/presentation/http/response"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
)

type CheckController struct {
	db *gorm.DB
}

func NewCheckController(db *gorm.DB) *CheckController {
	return &CheckController{
		db: db,
	}
}

// GET /api/v1/check
func (ctrl *CheckController) Check(c *fiber.Ctx) error {
	return response.Success(c, fiber.Map{
		"message": "API v1 is working!",
		"version": "1.0.0",
	})
}

// GET /api/v1/check/database
func (ctrl *CheckController) CheckDatabase(c *fiber.Ctx) error {
	sqlDB, err := ctrl.db.DB()
	if err != nil {
		return response.InternalServerError(c, "Failed to get database instance")
	}

	if err := sqlDB.Ping(); err != nil {
		return response.InternalServerError(c, "Database connection failed")
	}

	return response.Success(c, fiber.Map{
		"message": "Database connection successful",
	})
}

// GET /api/v1/check/broker
func (ctrl *CheckController) CheckBroker(c *fiber.Ctx) error {
	return response.Success(c, fiber.Map{
		"message": "RabbitMQ connection - TODO",
	})
}

// GET /api/v1/check/mailer
func (ctrl *CheckController) CheckMailer(c *fiber.Ctx) error {
	return response.Success(c, fiber.Map{
		"message": "Mailer service - TODO",
	})
}

// GET /api/v1/check/all
func (ctrl *CheckController) CheckAll(c *fiber.Ctx) error {
	return response.Success(c, fiber.Map{
		"database": "OK",
		"broker":   "TODO",
		"mailer":   "TODO",
	})
}
