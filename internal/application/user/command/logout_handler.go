package command

import (
	"apirest/internal/domain/master"
)

type LogoutCommand struct {
	Token  string
	UserID int64
	Role   string
}

type LogoutHandler struct {
	masterAccessLogRepo master.MasterAccessLogRepository
	// TODO: Add adminAccessLogRepo and employeeAccessLogRepo when ready
}

func NewLogoutHandler(
	masterAccessLogRepo master.MasterAccessLogRepository,
) *LogoutHandler {
	return &LogoutHandler{
		masterAccessLogRepo: masterAccessLogRepo,
	}
}

func (h *LogoutHandler) Handle(cmd LogoutCommand) error {
	// Invalidate the token
	return h.masterAccessLogRepo.InvalidateToken(cmd.Token)

	// TODO: When you have admin and employee repos, use role-based logic:
	/*
	switch cmd.Role {
	case "ROLE_MASTER":
		return h.masterAccessLogRepo.InvalidateToken(cmd.Token)
	case "ROLE_ADMIN":
		return h.adminAccessLogRepo.InvalidateToken(cmd.Token)
	case "ROLE_EMPLOYEE":
		return h.employeeAccessLogRepo.InvalidateToken(cmd.Token)
	default:
		return errors.NewValidationError("invalid role")
	}
	*/
}
