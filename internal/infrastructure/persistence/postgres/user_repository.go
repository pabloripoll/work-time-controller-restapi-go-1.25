package postgres

import (
	"apirest/internal/domain/shared/errors"
	"apirest/internal/domain/shared/valueobject"
	"apirest/internal/domain/user"
	"time"

	"gorm.io/gorm"
)

type PostgresUserRepository struct {
	db *gorm.DB
}

func NewPostgresUserRepository(db *gorm.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) Save(u *user.User) error {
	updatedAt := u.UpdatedAt()
	model := &UserModel{
		Role:            u.RoleString(),
		Email:           u.EmailString(),
		Password:        u.Password(),
		CreatedAt:       u.CreatedAt(),
		UpdatedAt:       &updatedAt,
		DeletedAt:       u.DeletedAt(),
		CreatedByUserID: 0,
	}

	return r.db.Create(model).Error
}

func (r *PostgresUserRepository) Update(u *user.User) error {
	updatedAt := u.UpdatedAt()
	updates := map[string]interface{}{
		"role":       u.RoleString(),
		"email":      u.EmailString(),
		"password":   u.Password(),
		"updated_at": updatedAt,
		"deleted_at": u.DeletedAt(),
	}

	return r.db.Model(&UserModel{}).
		Where("id = ?", u.ID()).
		Updates(updates).Error
}

func (r *PostgresUserRepository) FindByID(id int64) (*user.User, error) {
	var model UserModel
	err := r.db.Where("id = ?", id).First(&model).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewEntityNotFoundError("User", string(rune(id)))
		}
		return nil, err
	}

	return r.modelToDomain(&model)
}

func (r *PostgresUserRepository) FindByEmail(email valueobject.Email) (*user.User, error) {
	var model UserModel
	err := r.db.Where("email = ?", email.String()).First(&model).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewEntityNotFoundError("User", email.String())
		}
		return nil, err
	}

	return r.modelToDomain(&model)
}

func (r *PostgresUserRepository) FindAll() ([]*user.User, error) {
	var models []UserModel
	err := r.db.Find(&models).Error

	if err != nil {
		return nil, err
	}

	users := make([]*user.User, len(models))
	for i, model := range models {
		u, err := r.modelToDomain(&model)
		if err != nil {
			return nil, err
		}
		users[i] = u
	}

	return users, nil
}

func (r *PostgresUserRepository) Delete(id int64) error {
	// Soft delete
	now := time.Now()
	return r.db.Model(&UserModel{}).
		Where("id = ?", id).
		Update("deleted_at", &now).Error
}

func (r *PostgresUserRepository) ExistsByEmail(email valueobject.Email) (bool, error) {
	var count int64
	err := r.db.Model(&UserModel{}).
		Where("email = ?", email.String()).
		Where("deleted_at IS NULL").
		Count(&count).Error
	return count > 0, err
}

func (r *PostgresUserRepository) modelToDomain(model *UserModel) (*user.User, error) {
	email, err := valueobject.NewEmail(model.Email)
	if err != nil {
		return nil, err
	}

	role, err := valueobject.NewUserRole(model.Role)
	if err != nil {
		return nil, err
	}

	updatedAt := model.CreatedAt
	if model.UpdatedAt != nil {
		updatedAt = *model.UpdatedAt
	}

	return user.ReconstructUser(
		model.ID,
		email,
		model.Password,
		role,
		model.CreatedAt,
		updatedAt,
		model.DeletedAt,
	), nil
}
