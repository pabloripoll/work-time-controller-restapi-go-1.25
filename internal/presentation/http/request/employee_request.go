package request

import "time"

type CreateEmployeeRequest struct {
	Email       string     `json:"email" validate:"required,email"`
	Password    string     `json:"password" validate:"required,min=8"`
	Name        string     `json:"name" validate:"required"`
	Surname     string     `json:"surname" validate:"required"`
	PhoneNumber string     `json:"phone_number"`
	Department  string     `json:"department" validate:"required"`
	BirthDate   *time.Time `json:"birth_date,omitempty"`
}

type UpdateEmployeeProfileRequest struct {
	Name        string     `json:"name"`
	Surname     string     `json:"surname"`
	PhoneNumber string     `json:"phone_number"`
	Department  string     `json:"department"`
	BirthDate   *time.Time `json:"birth_date,omitempty"`
}
