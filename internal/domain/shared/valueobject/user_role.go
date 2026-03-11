package valueobject

import "apirest/internal/domain/shared/errors"

type UserRole struct {
	value string
}

const (
	roleMasterValue   = "ROLE_MASTER"
	roleAdminValue    = "ROLE_ADMIN"
	roleEmployeeValue = "ROLE_EMPLOYEE"
)

var (
	RoleMaster   = UserRole{value: roleMasterValue}
	RoleAdmin    = UserRole{value: roleAdminValue}
	RoleEmployee = UserRole{value: roleEmployeeValue}
)

func NewUserRole(role string) (UserRole, error) {
	switch role {
	case roleMasterValue:
		return RoleMaster, nil
	case roleAdminValue:
		return RoleAdmin, nil
	case roleEmployeeValue:
		return RoleEmployee, nil
	default:
		return UserRole{}, errors.NewValidationError("invalid user role: must be ROLE_MASTER, ROLE_ADMIN or ROLE_EMPLOYEE")
	}
}

func MustNewUserRole(role string) UserRole {
	r, err := NewUserRole(role)
	if err != nil {
		panic(err)
	}
	return r
}

func (r UserRole) String() string {
	return r.value
}

func (r UserRole) IsMaster() bool {
	return r.value == roleMasterValue
}

func (r UserRole) IsAdmin() bool {
	return r.value == roleAdminValue
}

func (r UserRole) IsEmployee() bool {
	return r.value == roleEmployeeValue
}

func (r UserRole) Equals(other UserRole) bool {
	return r.value == other.value
}

func (r UserRole) Validate() error {
	switch r.value {
	case roleMasterValue, roleAdminValue, roleEmployeeValue:
		return nil
	default:
		return errors.NewValidationError("invalid user role")
	}
}

func AllRoles() []UserRole {
	return []UserRole{RoleMaster, RoleAdmin, RoleEmployee}
}
