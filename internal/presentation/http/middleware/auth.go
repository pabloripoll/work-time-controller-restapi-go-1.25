package middleware

import (
	"apirest/internal/infrastructure/security/jwt"
	"apirest/internal/presentation/http/response"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type AuthMiddleware struct {
	jwtValidator *jwt.Validator
}

func NewAuthMiddleware(jwtValidator *jwt.Validator) *AuthMiddleware {
	return &AuthMiddleware{
		jwtValidator: jwtValidator,
	}
}

func (m *AuthMiddleware) Authenticate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return response.Unauthorized(c, "Missing authorization header")
		}

		// Extract token from "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return response.Unauthorized(c, "Invalid authorization header format")
		}

		token := parts[1]

		// Validate token
		claims, err := m.jwtValidator.ValidateToken(token)
		if err != nil {
			return response.Unauthorized(c, "Invalid or expired token")
		}

		// Store claims in context
		c.Locals("user_id", claims.UserID)
		c.Locals("email", claims.Email)
		c.Locals("role", claims.Role)
		c.Locals("token", token)

		return c.Next()
	}
}

func (m *AuthMiddleware) RequireRole(allowedRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		role := c.Locals("role").(string)

		for _, allowedRole := range allowedRoles {
			if role == allowedRole {
				return c.Next()
			}
		}

		return response.Forbidden(c, "Insufficient permissions")
	}
}

func (m *AuthMiddleware) RequireAdmin() fiber.Handler {
	return m.RequireRole("admin")
}

func (m *AuthMiddleware) RequireMaster() fiber.Handler {
	return m.RequireRole("master")
}

func (m *AuthMiddleware) RequireMember() fiber.Handler {
	return m.RequireRole("member")
}

func (m *AuthMiddleware) RequireAnyAuthenticated() fiber.Handler {
	return m.RequireRole("admin", "master", "member")
}
