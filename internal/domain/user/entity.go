package user

import (
	"apirest/internal/domain/shared/valueobject"
	"time"
)

type User struct {
	id        int64
	email     valueobject.Email
	password  string
	role      valueobject.UserRole
	createdAt time.Time
	updatedAt time.Time
	deletedAt *time.Time
}

func NewUser(
	email valueobject.Email,
	password string,
	role valueobject.UserRole,
) *User {
	now := time.Now()
	return &User{
		email:     email,
		password:  password,
		role:      role,
		createdAt: now,
		updatedAt: now,
		deletedAt: nil,
	}
}

func ReconstructUser(
	id int64,
	email valueobject.Email,
	password string,
	role valueobject.UserRole,
	createdAt, updatedAt time.Time,
	deletedAt *time.Time,
) *User {
	return &User{
		id:        id,
		email:     email,
		password:  password,
		role:      role,
		createdAt: createdAt,
		updatedAt: updatedAt,
		deletedAt: deletedAt,
	}
}

// Getters
func (u *User) ID() int64                      { return u.id }
func (u *User) Email() valueobject.Email       { return u.email }
func (u *User) EmailString() string            { return u.email.String() }
func (u *User) Password() string               { return u.password }
func (u *User) Role() valueobject.UserRole     { return u.role }
func (u *User) RoleString() string             { return u.role.String() }
func (u *User) CreatedAt() time.Time           { return u.createdAt }
func (u *User) UpdatedAt() time.Time           { return u.updatedAt }
func (u *User) DeletedAt() *time.Time          { return u.deletedAt }

// Role checks
func (u *User) IsMaster() bool {
	return u.role.IsMaster()
}

func (u *User) IsAdmin() bool {
	return u.role.IsAdmin()
}

func (u *User) IsEmployee() bool {
	return u.role.IsEmployee()
}

func (u *User) IsDeleted() bool {
	return u.deletedAt != nil
}

// Business methods
func (u *User) UpdatePassword(hashedPassword string) {
	u.password = hashedPassword
	u.updatedAt = time.Now()
}

func (u *User) UpdateEmail(email valueobject.Email) {
	u.email = email
	u.updatedAt = time.Now()
}

func (u *User) ChangeRole(role valueobject.UserRole) {
	u.role = role
	u.updatedAt = time.Now()
}

func (u *User) SoftDelete() {
	now := time.Now()
	u.deletedAt = &now
	u.updatedAt = now
}

func (u *User) Restore() {
	u.deletedAt = nil
	u.updatedAt = time.Now()
}
