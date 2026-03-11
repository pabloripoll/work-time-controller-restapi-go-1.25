package command

import (
	"apirest/internal/domain/master"
)

type UpdateMasterProfileCommand struct {
	MasterID int64
	Nickname string
	Avatar   *string
}

type UpdateMasterProfileHandler struct {
	profileRepo master.MasterProfileRepository
}

func NewUpdateMasterProfileHandler(
	profileRepo master.MasterProfileRepository,
) *UpdateMasterProfileHandler {
	return &UpdateMasterProfileHandler{
		profileRepo: profileRepo,
	}
}

func (h *UpdateMasterProfileHandler) Handle(cmd UpdateMasterProfileCommand) error {
	// Get existing profile
	profile, err := h.profileRepo.FindByMasterID(cmd.MasterID)
	if err != nil {
		return err
	}

	// Update fields
	if cmd.Nickname != "" {
		profile.UpdateNickname(cmd.Nickname)
	}

	if cmd.Avatar != nil {
		profile.UpdateAvatar(cmd.Avatar)
	}

	// Save changes
	return h.profileRepo.Update(profile)
}
