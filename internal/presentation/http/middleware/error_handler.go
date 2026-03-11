package middleware

import (
	"apirest/internal/presentation/http/response"
	"log"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()

		if err != nil {
			log.Printf("Error: %v", err)

			// Check if it's a Fiber error
			if e, ok := err.(*fiber.Error); ok {
				return response.Error(c, e.Code, e.Message, "ERROR")
			}

			// Default to internal server error
			return response.InternalServerError(c, err.Error())
		}

		return nil
	}
}
