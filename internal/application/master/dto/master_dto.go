package dto

import "time"

type MasterDTO struct {
	ID        int64          `json:"id"`
	Email     string         `json:"email"`
	Role      string         `json:"role"`
	Profile   *MasterProfile `json:"profile,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

type MasterFullDTO struct {
	Master  MasterSimpleDTO    `json:"master"`
	Profile MasterProfileDTO   `json:"profile"`
}

type MasterSimpleDTO struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
