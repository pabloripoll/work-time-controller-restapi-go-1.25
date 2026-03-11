package postgres

import (
	"apirest/internal/domain/master"
	"apirest/internal/domain/shared/errors"

	"gorm.io/gorm"
)

type PostgresMasterRepository struct {
	db *gorm.DB
}

func NewPostgresMasterRepository(db *gorm.DB) *PostgresMasterRepository {
	return &PostgresMasterRepository{db: db}
}

func (r *PostgresMasterRepository) Save(m *master.Master) error {
	model := &MasterModel{
		UserID:    m.UserID(),
		IsActive:  m.IsActive(),
		IsBanned:  m.IsBanned(),
		CreatedAt: m.CreatedAt(),
		UpdatedAt: m.UpdatedAt(),
	}

	result := r.db.Create(model)
	if result.Error != nil {
		return result.Error
	}

	// Update the ID in the domain entity using reflection or return new entity
	// For now, we'll need to fetch it back
	return nil
}

func (r *PostgresMasterRepository) FindByID(id int64) (*master.Master, error) {
	var model MasterModel
	err := r.db.Where("id = ?", id).First(&model).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewEntityNotFoundError("Master", string(rune(id)))
		}
		return nil, err
	}

	return r.modelToDomain(&model), nil
}

func (r *PostgresMasterRepository) FindByUserID(userID int64) (*master.Master, error) {
	var model MasterModel
	err := r.db.Where("user_id = ?", userID).First(&model).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewEntityNotFoundError("Master", string(rune(userID)))
		}
		return nil, err
	}

	return r.modelToDomain(&model), nil
}

func (r *PostgresMasterRepository) FindAll() ([]*master.Master, error) {
	var models []MasterModel
	err := r.db.Find(&models).Error

	if err != nil {
		return nil, err
	}

	masters := make([]*master.Master, len(models))
	for i, model := range models {
		masters[i] = r.modelToDomain(&model)
	}

	return masters, nil
}

func (r *PostgresMasterRepository) Delete(id int64) error {
	return r.db.Delete(&MasterModel{}, "id = ?", id).Error
}

func (r *PostgresMasterRepository) modelToDomain(model *MasterModel) *master.Master {
	return master.ReconstructMaster(
		model.ID,
		model.UserID,
		model.IsActive,
		model.IsBanned,
		model.CreatedAt,
		model.UpdatedAt,
	)
}
