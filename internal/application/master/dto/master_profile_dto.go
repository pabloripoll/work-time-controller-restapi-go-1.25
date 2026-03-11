package dto

import "time"

type MasterProfileDTO struct {
	ID        int64     `json:"id"`
	MasterID  int64     `json:"master_id"`
	Nickname  string    `json:"nickname"`
	Avatar    *string   `json:"avatar,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type MasterProfile struct {
	Nickname string  `json:"nickname"`
	Avatar   *string `json:"avatar,omitempty"`
}

type CreateMasterRequest struct {
	Email    string  `json:"email" validate:"required,email"`
	Password string  `json:"password" validate:"required,min=8"`
	Nickname string  `json:"nickname" validate:"required"`
	Avatar   *string `json:"avatar,omitempty"`
}

type UpdateMasterProfileRequest struct {
	Nickname string  `json:"nickname"`
	Avatar   *string `json:"avatar,omitempty"`
}
