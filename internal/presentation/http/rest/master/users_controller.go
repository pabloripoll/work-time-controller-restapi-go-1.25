package master

import (
	"apirest/internal/application/master/command"
	"apirest/internal/presentation/http/request"
	"apirest/internal/presentation/http/response"

	"github.com/gofiber/fiber/v2"
)

type UsersController struct {
	createMasterHandler *command.CreateMasterHandler
}

func NewUsersController(createMasterHandler *command.CreateMasterHandler) *UsersController {
	return &UsersController{
		createMasterHandler: createMasterHandler,
	}
}

// POST /api/v1/master/users
func (ctrl *UsersController) CreateMaster(c *fiber.Ctx) error {
	var req request.CreateMasterRequest

	if err := request.ParseAndValidate(c, &req); err != nil {
		return response.BadRequest(c, err.Error())
	}

	cmd := command.CreateMasterCommand{
		Email:    req.Email,
		Password: req.Password,
		Nickname: req.Nickname,
	}

	result, err := ctrl.createMasterHandler.Handle(cmd)
	if err != nil {
		return response.HandleDomainError(c, err)
	}

	return response.Created(c, result)
}

// GET /api/v1/master/users
func (ctrl *UsersController) GetAllMasters(c *fiber.Ctx) error {
	// TODO: Implement
	return response.Success(c, fiber.Map{"message": "Get all masters - TODO"})
}

// GET /api/v1/master/users/:id/profiles
func (ctrl *UsersController) GetMasterProfile(c *fiber.Ctx) error {
	// TODO: Implement
	return response.Success(c, fiber.Map{"message": "Get master profile - TODO"})
}

// DELETE /api/v1/master/users/:id
func (ctrl *UsersController) DeleteMaster(c *fiber.Ctx) error {
	// TODO: Implement
	return response.NoContent(c)
}
