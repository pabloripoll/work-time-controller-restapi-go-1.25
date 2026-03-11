package master

import (
	"time"
)

type Master struct {
	id        int64
	userID    int64
	isActive  bool
	isBanned  bool
	createdAt time.Time
	updatedAt time.Time
}

func NewMaster(userID int64) *Master {
	now := time.Now()
	return &Master{
		userID:    userID,
		isActive:  true,
		isBanned:  false,
		createdAt: now,
		updatedAt: now,
	}
}

func ReconstructMaster(
	id int64,
	userID int64,
	isActive, isBanned bool,
	createdAt, updatedAt time.Time,
) *Master {
	return &Master{
		id:        id,
		userID:    userID,
		isActive:  isActive,
		isBanned:  isBanned,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

// Getters
func (m *Master) ID() int64           { return m.id }
func (m *Master) UserID() int64       { return m.userID }
func (m *Master) IsActive() bool      { return m.isActive }
func (m *Master) IsBanned() bool      { return m.isBanned }
func (m *Master) CreatedAt() time.Time { return m.createdAt }
func (m *Master) UpdatedAt() time.Time { return m.updatedAt }

// Business methods
func (m *Master) Activate() {
	m.isActive = true
	m.updatedAt = time.Now()
}

func (m *Master) Deactivate() {
	m.isActive = false
	m.updatedAt = time.Now()
}

func (m *Master) Ban() {
	m.isBanned = true
	m.updatedAt = time.Now()
}

func (m *Master) Unban() {
	m.isBanned = false
	m.updatedAt = time.Now()
}
