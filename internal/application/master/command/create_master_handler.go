package command

import (
	"apirest/internal/domain/master"
	"apirest/internal/domain/shared/valueobject"
	"apirest/internal/domain/user"
)

type CreateMasterCommand struct {
	Email    string
	Password string
	Nickname string
}

type CreateMasterHandler struct {
	userService *user.Service
	masterRepo  master.MasterRepository
	profileRepo master.MasterProfileRepository
}

func NewCreateMasterHandler(
	userService *user.Service,
	masterRepo master.MasterRepository,
	profileRepo master.MasterProfileRepository,
) *CreateMasterHandler {
	return &CreateMasterHandler{
		userService: userService,
		masterRepo:  masterRepo,
		profileRepo: profileRepo,
	}
}

func (h *CreateMasterHandler) Handle(cmd CreateMasterCommand) (*master.Master, error) {
	// Validate email
	email, err := valueobject.NewEmail(cmd.Email)
	if err != nil {
		return nil, err
	}

	// Create user
	userEntity, err := h.userService.CreateMaster(email, cmd.Password)
	if err != nil {
		return nil, err
	}

	// Create master
	masterEntity := master.NewMaster(userEntity.ID())
	if err := h.masterRepo.Save(masterEntity); err != nil {
		return nil, err
	}

	// Create profile
	profile := master.NewMasterProfile(masterEntity.ID(), cmd.Nickname)
	if err := h.profileRepo.Save(profile); err != nil {
		return nil, err
	}

	return masterEntity, nil
}
