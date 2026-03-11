package query

import (
	"apirest/internal/application/user/dto"
	"apirest/internal/domain/user"
)

type WhoAmIQuery struct {
	UserID int64
}

type WhoAmIHandler struct {
	repoUser user.UserRepository
}

func NewWhoAmIHandler(repoUser user.UserRepository) *WhoAmIHandler {
	return &WhoAmIHandler{
		repoUser: repoUser,
	}
}

func (h *WhoAmIHandler) Handle(query WhoAmIQuery) (*dto.UserDTO, error) {
	u, err := h.repoUser.FindByID(query.UserID)
	if err != nil {
		return nil, err
	}

	return &dto.UserDTO{
		ID:        u.ID(),
		Email:     u.EmailString(),
		Role:      u.RoleString(),
		CreatedAt: u.CreatedAt(),
		UpdatedAt: u.UpdatedAt(),
		DeletedAt: u.DeletedAt(),
	}, nil
}
