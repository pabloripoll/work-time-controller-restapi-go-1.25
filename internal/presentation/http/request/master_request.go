package request

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

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}
