package postgres

import (
	"apirest/internal/domain/master"
	"apirest/internal/domain/shared/errors"

	"gorm.io/gorm"
)

type PostgresMasterProfileRepository struct {
	db *gorm.DB
}

func NewPostgresMasterProfileRepository(db *gorm.DB) *PostgresMasterProfileRepository {
	return &PostgresMasterProfileRepository{db: db}
}

func (r *PostgresMasterProfileRepository) Save(profile *master.MasterProfile) error {
	model := &MasterProfileModel{
		MasterID:  profile.MasterID(),
		Nickname:  profile.Nickname(),
		Avatar:    profile.Avatar(),
		CreatedAt: profile.CreatedAt(),
		UpdatedAt: profile.UpdatedAt(),
	}

	return r.db.Create(model).Error
}

func (r *PostgresMasterProfileRepository) Update(profile *master.MasterProfile) error {
	updates := map[string]interface{}{
		"nickname":   profile.Nickname(),
		"avatar":     profile.Avatar(),
		"updated_at": profile.UpdatedAt(),
	}

	return r.db.Model(&MasterProfileModel{}).
		Where("id = ?", profile.ID()).
		Updates(updates).Error
}

func (r *PostgresMasterProfileRepository) FindByID(id int64) (*master.MasterProfile, error) {
	var model MasterProfileModel
	err := r.db.Where("id = ?", id).First(&model).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewEntityNotFoundError("MasterProfile", string(rune(id)))
		}
		return nil, err
	}

	return r.modelToDomain(&model), nil
}

func (r *PostgresMasterProfileRepository) FindByMasterID(masterID int64) (*master.MasterProfile, error) {
	var model MasterProfileModel
	err := r.db.Where("master_id = ?", masterID).First(&model).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewEntityNotFoundError("MasterProfile", string(rune(masterID)))
		}
		return nil, err
	}

	return r.modelToDomain(&model), nil
}

func (r *PostgresMasterProfileRepository) Delete(id int64) error {
	return r.db.Delete(&MasterProfileModel{}, "id = ?", id).Error
}

func (r *PostgresMasterProfileRepository) modelToDomain(model *MasterProfileModel) *master.MasterProfile {
	return master.ReconstructMasterProfile(
		model.ID,
		model.MasterID,
		model.Nickname,
		model.Avatar,
		model.CreatedAt,
		model.UpdatedAt,
	)
}
