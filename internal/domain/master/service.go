package master

type MasterService struct {
	repoProfile   MasterProfileRepository
	repoAccessLog MasterAccessLogRepository
}

func NewMasterService(repoProfile MasterProfileRepository, repoAccessLog MasterAccessLogRepository) *MasterService {
	return &MasterService{
		repoProfile:   repoProfile,
		repoAccessLog: repoAccessLog,
	}
}

func (s *MasterService) CreateProfile(profile *MasterProfile) error {
	return s.repoProfile.Save(profile)
}

func (s *MasterService) UpdateProfile(profile *MasterProfile) error {
	return s.repoProfile.Update(profile)
}

func (s *MasterService) GetProfileByMasterID(masterID int64) (*MasterProfile, error) {
	return s.repoProfile.FindByMasterID(masterID)
}

func (s *MasterService) CreateAccessLog(log *MasterAccessLog) error {
	return s.repoAccessLog.Save(log)
}

func (s *MasterService) TerminateAccessLog(logID int64) error {
	log, err := s.repoAccessLog.FindByID(logID)
	if err != nil {
		return err
	}

	log.Terminate()
	return s.repoAccessLog.Update(log)
}

func (s *MasterService) TerminateAllUserAccessLogs(userID int64) error {
	return s.repoAccessLog.TerminateAllByUserID(userID)
}

func (s *MasterService) IncrementRequestCount(token string) error {
	log, err := s.repoAccessLog.FindByToken(token)
	if err != nil {
		return err
	}

	log.IncrementRequestCount()
	log.CheckExpiration()

	return s.repoAccessLog.Update(log)
}
