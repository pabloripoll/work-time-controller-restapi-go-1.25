package master

import (
	"apirest/internal/application/master/command"
	"apirest/internal/application/master/query"
	"apirest/internal/presentation/http/request"
	"apirest/internal/presentation/http/response"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type AccountController struct {
	getProfileHandler    *query.GetMasterProfileHandler
	updateProfileHandler *command.UpdateMasterProfileHandler
}

func NewAccountController(
	getProfileHandler *query.GetMasterProfileHandler,
	updateProfileHandler *command.UpdateMasterProfileHandler,
) *AccountController {
	return &AccountController{
		getProfileHandler:    getProfileHandler,
		updateProfileHandler: updateProfileHandler,
	}
}

// GET /api/v1/master/account/profile
func (ctrl *AccountController) GetProfile(c *fiber.Ctx) error {
	masterIDStr := c.Locals("master_id").(string) // Middleware should set this

	masterID, err := strconv.ParseInt(masterIDStr, 10, 64)
	if err != nil {
		return response.BadRequest(c, "invalid master ID")
	}

	q := query.GetMasterProfileQuery{
		MasterID: masterID,
	}

	result, err := ctrl.getProfileHandler.Handle(q)
	if err != nil {
		return response.HandleDomainError(c, err)
	}

	return response.Success(c, result)
}

// PATCH /api/v1/master/account/profile
func (ctrl *AccountController) UpdateProfile(c *fiber.Ctx) error {
	masterIDStr := c.Locals("master_id").(string)

	masterID, err := strconv.ParseInt(masterIDStr, 10, 64)
	if err != nil {
		return response.BadRequest(c, "invalid master ID")
	}

	var req request.UpdateMasterProfileRequest
	if err := request.ParseAndValidate(c, &req); err != nil {
		return response.BadRequest(c, err.Error())
	}

	cmd := command.UpdateMasterProfileCommand{
		MasterID: masterID,
		Nickname: req.Nickname, // Fixed typo
		Avatar:   req.Avatar,
	}

	err = ctrl.updateProfileHandler.Handle(cmd)
	if err != nil {
		return response.HandleDomainError(c, err)
	}

	return response.Success(c, fiber.Map{"message": "Profile updated successfully"})
}
