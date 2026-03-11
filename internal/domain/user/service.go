package user

import (
	"apirest/internal/domain/shared/errors"
	"apirest/internal/domain/shared/valueobject"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo UserRepository
}

func NewService(repo UserRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateMaster(email valueobject.Email, password string) (*User, error) {
	exists, err := s.repo.ExistsByEmail(email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.NewDuplicateError("email", email.String())
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := NewUser(email, string(hashedPassword), valueobject.RoleMaster)
	if err := s.repo.Save(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) CreateAdmin(email valueobject.Email, password string) (*User, error) {
	exists, err := s.repo.ExistsByEmail(email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.NewDuplicateError("email", email.String())
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := NewUser(email, string(hashedPassword), valueobject.RoleAdmin)
	if err := s.repo.Save(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) CreateEmployee(email valueobject.Email, password string) (*User, error) {
	exists, err := s.repo.ExistsByEmail(email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.NewDuplicateError("email", email.String())
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := NewUser(email, string(hashedPassword), valueobject.RoleEmployee)
	if err := s.repo.Save(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) Authenticate(email valueobject.Email, password string) (*User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, errors.NewInvalidCredentialsError()
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password()), []byte(password)); err != nil {
		return nil, errors.NewInvalidCredentialsError()
	}

	return user, nil
}

func (s *Service) GetByID(id int64) (*User, error) {
	return s.repo.FindByID(id)
}

func (s *Service) GetByEmail(email valueobject.Email) (*User, error) {
	return s.repo.FindByEmail(email)
}

func (s *Service) UpdatePassword(userID int64, newPassword string) error {
	user, err := s.repo.FindByID(userID)
	if err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.UpdatePassword(string(hashedPassword))
	return s.repo.Update(user)
}
