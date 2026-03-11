package command

import (
	"apirest/internal/application/user/dto"
	"apirest/internal/domain/master"
	"apirest/internal/domain/shared/errors"
	"apirest/internal/domain/shared/valueobject"
	"apirest/internal/domain/user"
	"apirest/internal/infrastructure/security/jwt"
)

type LoginMasterCommand struct {
	Email     string
	Password  string
	IPAddress string
	UserAgent string
}

type LoginMasterHandler struct {
	userService   *user.Service
	masterRepo    master.MasterRepository
	accessLogRepo master.MasterAccessLogRepository
	jwtGenerator  *jwt.Generator
}

func NewLoginMasterHandler(
	userService *user.Service,
	masterRepo master.MasterRepository,
	accessLogRepo master.MasterAccessLogRepository,
	jwtGenerator *jwt.Generator,
) *LoginMasterHandler {
	return &LoginMasterHandler{
		userService:   userService,
		masterRepo:    masterRepo,
		accessLogRepo: accessLogRepo,
		jwtGenerator:  jwtGenerator,
	}
}

func (h *LoginMasterHandler) Handle(cmd LoginMasterCommand) (*dto.LoginResponse, error) {
	// Parse email
	email, err := valueobject.NewEmail(cmd.Email)
	if err != nil {
		return nil, err
	}

	// Authenticate user
	u, err := h.userService.Authenticate(email, cmd.Password)
	if err != nil {
		return nil, err
	}

	// Check if user is master
	if !u.IsMaster() {
		return nil, errors.NewForbiddenError("user is not a master")
	}

	// Get master entity
	masterEntity, err := h.masterRepo.FindByUserID(u.ID())
	if err != nil {
		return nil, err
	}

	// Generate tokens
	accessToken, accessExpiresAt, err := h.jwtGenerator.GenerateAccessToken(u.ID(), u.EmailString(), u.RoleString())
	if err != nil {
		return nil, err
	}

	refreshToken, refreshExpiresAt, err := h.jwtGenerator.GenerateRefreshToken(u.ID(), u.EmailString(), u.RoleString())
	if err != nil {
		return nil, err
	}

	// Create access log
	accessLog := master.NewMasterAccessLog(
		masterEntity.ID(),
		u.ID(),
		refreshToken,
		refreshExpiresAt,
		cmd.IPAddress,
		cmd.UserAgent,
		nil,
	)

	if err := h.accessLogRepo.Save(accessLog); err != nil {
		return nil, err
	}

	// Build response
	return &dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    accessExpiresAt.Unix(),
		User: dto.UserDTO{
			ID:        u.ID(),
			Email:     u.EmailString(),
			Role:      u.RoleString(),
			CreatedAt: u.CreatedAt(),
			UpdatedAt: u.UpdatedAt(),
			DeletedAt: u.DeletedAt(),
		},
	}, nil
}
