package response

import (
	"apirest/internal/domain/shared/errors"
	"github.com/gofiber/fiber/v2"
)

func HandleDomainError(c *fiber.Ctx, err error) error {
	switch e := err.(type) {
	case *errors.EntityNotFoundError:
		return NotFound(c, e.Message)
	case *errors.ValidationError:
		return BadRequest(c, e.Message)
	case *errors.InvalidEmailError:
		return BadRequest(c, e.Message)
	case *errors.InvalidUUIDError:
		return BadRequest(c, e.Message)
	case *errors.UnauthorizedError:
		return Unauthorized(c, e.Message)
	case *errors.ForbiddenError:
		return Forbidden(c, e.Message)
	case *errors.DuplicateError:
		return Error(c, fiber.StatusConflict, e.Message, "DUPLICATE_ENTRY")
	case *errors.InvalidCredentialsError:
		return Unauthorized(c, e.Message)
	default:
		return InternalServerError(c, "An unexpected error occurred")
	}
}
