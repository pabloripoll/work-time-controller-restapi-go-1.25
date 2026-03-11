package master

type MasterRepository interface {
	Save(master *Master) error
	FindByID(id int64) (*Master, error)
	FindByUserID(userID int64) (*Master, error)
	FindAll() ([]*Master, error)
	Delete(id int64) error
}

type MasterProfileRepository interface {
	Save(profile *MasterProfile) error
	Update(profile *MasterProfile) error
	FindByID(id int64) (*MasterProfile, error)
	FindByMasterID(masterID int64) (*MasterProfile, error)
	Delete(id int64) error
}

type MasterAccessLogRepository interface {
	Save(log *MasterAccessLog) error
	Update(log *MasterAccessLog) error
	FindByID(id int64) (*MasterAccessLog, error)
	FindByToken(token string) (*MasterAccessLog, error)
	FindByMasterID(masterID int64) ([]*MasterAccessLog, error)
	FindByUserID(userID int64) ([]*MasterAccessLog, error)
	FindActiveByUserID(userID int64) ([]*MasterAccessLog, error)
	InvalidateToken(token string) error
	TerminateAllByUserID(userID int64) error
}
