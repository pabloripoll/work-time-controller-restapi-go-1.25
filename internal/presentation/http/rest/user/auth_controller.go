package user

import (
	"apirest/internal/application/user/command"
	"apirest/internal/application/user/query"
	"apirest/internal/presentation/http/request"
	"apirest/internal/presentation/http/response"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	refreshTokenHandler *command.RefreshTokenHandler
	logoutHandler       *command.LogoutHandler
	whoAmIHandler       *query.WhoAmIHandler
}

func NewAuthController(
	refreshTokenHandler *command.RefreshTokenHandler,
	logoutHandler *command.LogoutHandler,
	whoAmIHandler *query.WhoAmIHandler,
) *AuthController {
	return &AuthController{
		refreshTokenHandler: refreshTokenHandler,
		logoutHandler:       logoutHandler,
		whoAmIHandler:       whoAmIHandler,
	}
}

// POST /api/v1/auth/refresh
func (ctrl *AuthController) RefreshToken(c *fiber.Ctx) error {
	var req struct {
		RefreshToken string `json:"refresh_token" validate:"required"`
	}

	if err := request.ParseAndValidate(c, &req); err != nil {
		return response.BadRequest(c, err.Error())
	}

	cmd := command.RefreshTokenCommand{
		RefreshToken: req.RefreshToken,
	}

	result, err := ctrl.refreshTokenHandler.Handle(cmd)
	if err != nil {
		return response.HandleDomainError(c, err)
	}

	return response.Success(c, result)
}

// POST /api/v1/auth/logout
func (ctrl *AuthController) Logout(c *fiber.Ctx) error {
	userIDStr := c.Locals("user_id").(string)
	role := c.Locals("role").(string)
	token := c.Locals("token").(string)

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		return response.BadRequest(c, "invalid user ID")
	}

	cmd := command.LogoutCommand{
		Token:  token,
		UserID: userID,
		Role:   role,
	}

	if err := ctrl.logoutHandler.Handle(cmd); err != nil {
		return response.HandleDomainError(c, err)
	}

	return response.Success(c, fiber.Map{"message": "Logged out successfully"})
}

// GET /api/v1/auth/whoami
func (ctrl *AuthController) WhoAmI(c *fiber.Ctx) error {
	userIDStr := c.Locals("user_id").(string)

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		return response.BadRequest(c, "invalid user ID")
	}

	q := query.WhoAmIQuery{
		UserID: userID,
	}

	result, err := ctrl.whoAmIHandler.Handle(q)
	if err != nil {
		return response.HandleDomainError(c, err)
	}

	return response.Success(c, result)
}
