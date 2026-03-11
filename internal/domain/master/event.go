package master

import (
	//"apirest/internal/domain/shared/valueobject" -> valueobject.UUID instead of int64
	"time"
)

type MasterCreatedEvent struct {
	UserID     int64
	Email      string
	Name       string
	Surname    string
	Department string
	OccurredAt time.Time
}

func NewMasterCreatedEvent(userID int64, email, name, surname, department string) MasterCreatedEvent {
	return MasterCreatedEvent{
		UserID:     userID,
		Email:      email,
		Name:       name,
		Surname:    surname,
		Department: department,
		OccurredAt: time.Now().UTC(),
	}
}

type MasterLoggedInEvent struct {
	UserID     int64
	Email      string
	IPAddress  string
	UserAgent  string
	OccurredAt time.Time
}

func NewMasterLoggedInEvent(userID int64, email, ipAddress, userAgent string) MasterLoggedInEvent {
	return MasterLoggedInEvent{
		UserID:     userID,
		Email:      email,
		IPAddress:  ipAddress,
		UserAgent:  userAgent,
		OccurredAt: time.Now().UTC(),
	}
}
