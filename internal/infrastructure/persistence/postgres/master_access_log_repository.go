package postgres

import (
	"apirest/internal/domain/master"
	"apirest/internal/domain/shared/errors"

	"gorm.io/gorm"
)

type PostgresMasterAccessLogRepository struct {
	db *gorm.DB
}

func NewPostgresMasterAccessLogRepository(db *gorm.DB) *PostgresMasterAccessLogRepository {
	return &PostgresMasterAccessLogRepository{db: db}
}

func (r *PostgresMasterAccessLogRepository) Save(log *master.MasterAccessLog) error {
	payloadStr := string(log.PayloadJSON())
	ipAddr := log.IPAddress()
	userAgent := log.UserAgent()

	model := &MasterAccessLogModel{
		MasterID:      log.MasterID(),
		UserID:        log.UserID(),
		Token:         log.Token(),
		IsTerminated:  log.IsTerminated(),
		IsExpired:     log.IsExpired(),
		ExpiresAt:     log.ExpiresAt(),
		RefreshCount:  log.RefreshCount(),
		CreatedAt:     log.CreatedAt(),
		UpdatedAt:     log.UpdatedAt(),
		IPAddress:     &ipAddr,
		UserAgent:     &userAgent,
		RequestsCount: log.RequestsCount(),
		Payload:       &payloadStr,
	}

	return r.db.Create(model).Error
}

func (r *PostgresMasterAccessLogRepository) Update(log *master.MasterAccessLog) error {
	updates := map[string]interface{}{
		"is_terminated":  log.IsTerminated(),
		"is_expired":     log.IsExpired(),
		"refresh_count":  log.RefreshCount(),
		"updated_at":     log.UpdatedAt(),
		"requests_count": log.RequestsCount(),
	}

	return r.db.Model(&MasterAccessLogModel{}).
		Where("id = ?", log.ID()).
		Updates(updates).Error
}

func (r *PostgresMasterAccessLogRepository) FindByID(id int64) (*master.MasterAccessLog, error) {
	var model MasterAccessLogModel
	err := r.db.Where("id = ?", id).First(&model).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewEntityNotFoundError("MasterAccessLog", string(rune(id)))
		}
		return nil, err
	}

	return r.modelToDomain(&model), nil
}

func (r *PostgresMasterAccessLogRepository) FindByToken(token string) (*master.MasterAccessLog, error) {
	var model MasterAccessLogModel
	err := r.db.Where("token = ?", token).First(&model).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewEntityNotFoundError("MasterAccessLog", "token")
		}
		return nil, err
	}

	return r.modelToDomain(&model), nil
}

func (r *PostgresMasterAccessLogRepository) FindByMasterID(masterID int64) ([]*master.MasterAccessLog, error) {
	var models []MasterAccessLogModel
	err := r.db.Where("master_id = ?", masterID).
		Order("created_at DESC").
		Find(&models).Error

	if err != nil {
		return nil, err
	}

	logs := make([]*master.MasterAccessLog, len(models))
	for i, model := range models {
		logs[i] = r.modelToDomain(&model)
	}

	return logs, nil
}

func (r *PostgresMasterAccessLogRepository) FindByUserID(userID int64) ([]*master.MasterAccessLog, error) {
	var models []MasterAccessLogModel
	err := r.db.Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&models).Error

	if err != nil {
		return nil, err
	}

	logs := make([]*master.MasterAccessLog, len(models))
	for i, model := range models {
		logs[i] = r.modelToDomain(&model)
	}

	return logs, nil
}

func (r *PostgresMasterAccessLogRepository) FindActiveByUserID(userID int64) ([]*master.MasterAccessLog, error) {
	var models []MasterAccessLogModel
	err := r.db.Where("user_id = ? AND is_terminated = ? AND is_expired = ?", userID, false, false).
		Find(&models).Error

	if err != nil {
		return nil, err
	}

	logs := make([]*master.MasterAccessLog, len(models))
	for i, model := range models {
		logs[i] = r.modelToDomain(&model)
	}

	return logs, nil
}

func (r *PostgresMasterAccessLogRepository) InvalidateToken(token string) error {
	return r.db.Model(&MasterAccessLogModel{}).
		Where("token = ?", token).
		Updates(map[string]interface{}{
			"is_terminated": true,
		}).Error
}

func (r *PostgresMasterAccessLogRepository) TerminateAllByUserID(userID int64) error {
	return r.db.Model(&MasterAccessLogModel{}).
		Where("user_id = ? AND is_terminated = ?", userID, false).
		Update("is_terminated", true).Error
}

func (r *PostgresMasterAccessLogRepository) modelToDomain(model *MasterAccessLogModel) *master.MasterAccessLog {
	var payloadBytes []byte
	if model.Payload != nil {
		payloadBytes = []byte(*model.Payload)
	}

	ipAddress := ""
	if model.IPAddress != nil {
		ipAddress = *model.IPAddress
	}

	userAgent := ""
	if model.UserAgent != nil {
		userAgent = *model.UserAgent
	}

	return master.ReconstructMasterAccessLog(
		model.ID,
		model.MasterID,
		model.UserID,
		model.Token,
		model.IsTerminated,
		model.IsExpired,
		model.ExpiresAt,
		model.CreatedAt,
		model.UpdatedAt,
		model.RefreshCount,
		ipAddress,
		userAgent,
		model.RequestsCount,
		payloadBytes,
	)
}
