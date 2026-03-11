package valueobject

import (
	"apirest/internal/domain/shared/errors"
	"regexp"
	"strings"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

type Email struct {
	value string
}

func NewEmail(email string) (Email, error) {
	email = strings.TrimSpace(strings.ToLower(email))

	if email == "" {
		return Email{}, errors.NewValidationError("email cannot be empty")
	}

	if len(email) > 64 {
		return Email{}, errors.NewValidationError("email cannot exceed 64 characters")
	}

	if !emailRegex.MatchString(email) {
		return Email{}, errors.NewInvalidEmailError(email)
	}

	return Email{value: email}, nil
}

func MustNewEmail(email string) Email {
	e, err := NewEmail(email)
	if err != nil {
		panic(err)
	}
	return e
}

func (e Email) String() string {
	return e.value
}

func (e Email) Equals(other Email) bool {
	return e.value == other.value
}
