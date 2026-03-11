package master

import (
	"apirest/internal/application/master/command"
	"apirest/internal/presentation/http/request"
	"apirest/internal/presentation/http/response"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	loginHandler *command.LoginMasterHandler
}

func NewAuthController(loginHandler *command.LoginMasterHandler) *AuthController {
	return &AuthController{
		loginHandler: loginHandler,
	}
}

// POST /api/v1/master/auth/login
func (ctrl *AuthController) Login(c *fiber.Ctx) error {
	var req request.LoginRequest

	if err := request.ParseAndValidate(c, &req); err != nil {
		return response.BadRequest(c, err.Error())
	}

	cmd := command.LoginMasterCommand{
		Email:     req.Email,
		Password:  req.Password,
		IPAddress: c.IP(),
		UserAgent: c.Get("User-Agent"),
	}

	result, err := ctrl.loginHandler.Handle(cmd)
	if err != nil {
		return response.HandleDomainError(c, err)
	}

	return response.Success(c, result)
}
