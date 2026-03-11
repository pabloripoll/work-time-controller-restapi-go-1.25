package user

import (
	//"apirest/internal/domain/shared/valueobject" -> valueobject.UUID instead of int64
	"time"
)

// Domain Events
type UserCreatedEvent struct {
	UserID      int64
	Email       string
	Role        string
	OccurredAt  time.Time
}

func NewUserCreatedEvent(user *User) UserCreatedEvent {
	return UserCreatedEvent{
		UserID:     user.ID(),
		Email:      user.Email().String(),
		Role:       user.Role().String(),
		OccurredAt: time.Now().UTC(),
	}
}

type UserLoggedInEvent struct {
	UserID     int64
	Email      string
	IPAddress  string
	UserAgent  string
	OccurredAt time.Time
}

func NewUserLoggedInEvent(user *User, ipAddress, userAgent string) UserLoggedInEvent {
	return UserLoggedInEvent{
		UserID:     user.ID(),
		Email:      user.Email().String(),
		IPAddress:  ipAddress,
		UserAgent:  userAgent,
		OccurredAt: time.Now().UTC(),
	}
}

type UserDeletedEvent struct {
	UserID     int64
	Email      string
	OccurredAt time.Time
}

func NewUserDeletedEvent(user *User) UserDeletedEvent {
	return UserDeletedEvent{
		UserID:     user.ID(),
		Email:      user.Email().String(),
		OccurredAt: time.Now().UTC(),
	}
}
