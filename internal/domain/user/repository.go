package user

import (
	"apirest/internal/domain/shared/valueobject"
)

// UserRepository defines the contract for user persistence
type UserRepository interface {
	Save(user *User) error
	Update(user *User) error
	FindByID(id int64) (*User, error)
	FindByEmail(email valueobject.Email) (*User, error)
	FindAll() ([]*User, error)
	Delete(id int64) error
	ExistsByEmail(email valueobject.Email) (bool, error)
}