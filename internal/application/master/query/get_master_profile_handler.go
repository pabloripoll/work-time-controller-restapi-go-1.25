package query

import (
	"apirest/internal/application/master/dto"
	"apirest/internal/domain/master"
	"apirest/internal/domain/user"
)

type GetMasterProfileQuery struct {
	MasterID int64
}

type GetMasterProfileHandler struct {
	masterRepo  master.MasterRepository
	profileRepo master.MasterProfileRepository
	userRepo    user.UserRepository
}

func NewGetMasterProfileHandler(
	masterRepo master.MasterRepository,
	profileRepo master.MasterProfileRepository,
	userRepo user.UserRepository,
) *GetMasterProfileHandler {
	return &GetMasterProfileHandler{
		masterRepo:  masterRepo,
		profileRepo: profileRepo,
		userRepo:    userRepo,
	}
}

func (h *GetMasterProfileHandler) Handle(query GetMasterProfileQuery) (*dto.MasterFullDTO, error) {
	// Get master
	masterEntity, err := h.masterRepo.FindByID(query.MasterID)
	if err != nil {
		return nil, err
	}

	// Get user
	userEntity, err := h.userRepo.FindByID(masterEntity.UserID())
	if err != nil {
		return nil, err
	}

	// Get profile
	profile, err := h.profileRepo.FindByMasterID(query.MasterID)
	if err != nil {
		return nil, err
	}

	return &dto.MasterFullDTO{
		Master: dto.MasterSimpleDTO{
			ID:        masterEntity.ID(),
			Email:     userEntity.EmailString(),
			CreatedAt: masterEntity.CreatedAt(),
			UpdatedAt: masterEntity.UpdatedAt(),
		},
		Profile: dto.MasterProfileDTO{
			ID:        profile.ID(),
			MasterID:  profile.MasterID(),
			Nickname:  profile.Nickname(),
			Avatar:    profile.Avatar(),
			CreatedAt: profile.CreatedAt(),
			UpdatedAt: profile.UpdatedAt(),
		},
	}, nil
}
