package master

import (
	"encoding/json"
	"time"
)

type MasterAccessLog struct {
	id            int64
	masterID      int64
	userID        int64
	token         string
	isTerminated  bool
	isExpired     bool
	expiresAt     time.Time
	refreshCount  int
	createdAt     time.Time
	updatedAt     time.Time
	ipAddress     string
	userAgent     string
	requestsCount int
	payload       map[string]interface{}
}

func NewMasterAccessLog(
	masterID int64,
	userID int64,
	token string,
	expiresAt time.Time,
	ipAddress, userAgent string,
	payload map[string]interface{},
) *MasterAccessLog {
	now := time.Now()

	return &MasterAccessLog{
		masterID:      masterID,
		userID:        userID,
		token:         token,
		isTerminated:  false,
		isExpired:     false,
		expiresAt:     expiresAt,
		refreshCount:  0,
		createdAt:     now,
		updatedAt:     now,
		ipAddress:     ipAddress,
		userAgent:     userAgent,
		requestsCount: 0,
		payload:       payload,
	}
}

func ReconstructMasterAccessLog(
	id int64,
	masterID int64,
	userID int64,
	token string,
	isTerminated, isExpired bool,
	expiresAt, createdAt, updatedAt time.Time,
	refreshCount int,
	ipAddress, userAgent string,
	requestsCount int,
	payloadJSON []byte,
) *MasterAccessLog {
	var payload map[string]interface{}
	if len(payloadJSON) > 0 {
		json.Unmarshal(payloadJSON, &payload)
	}

	return &MasterAccessLog{
		id:            id,
		masterID:      masterID,
		userID:        userID,
		token:         token,
		isTerminated:  isTerminated,
		isExpired:     isExpired,
		expiresAt:     expiresAt,
		refreshCount:  refreshCount,
		createdAt:     createdAt,
		updatedAt:     updatedAt,
		ipAddress:     ipAddress,
		userAgent:     userAgent,
		requestsCount: requestsCount,
		payload:       payload,
	}
}

// Getters
func (al *MasterAccessLog) ID() int64                      { return al.id }
func (al *MasterAccessLog) MasterID() int64                { return al.masterID }
func (al *MasterAccessLog) UserID() int64                  { return al.userID }
func (al *MasterAccessLog) Token() string                  { return al.token }
func (al *MasterAccessLog) IsTerminated() bool             { return al.isTerminated }
func (al *MasterAccessLog) IsExpired() bool                { return al.isExpired }
func (al *MasterAccessLog) ExpiresAt() time.Time           { return al.expiresAt }
func (al *MasterAccessLog) RefreshCount() int              { return al.refreshCount }
func (al *MasterAccessLog) CreatedAt() time.Time           { return al.createdAt }
func (al *MasterAccessLog) UpdatedAt() time.Time           { return al.updatedAt }
func (al *MasterAccessLog) IPAddress() string              { return al.ipAddress }
func (al *MasterAccessLog) UserAgent() string              { return al.userAgent }
func (al *MasterAccessLog) RequestsCount() int             { return al.requestsCount }
func (al *MasterAccessLog) Payload() map[string]interface{} { return al.payload }

// Business methods
func (al *MasterAccessLog) Terminate() {
	al.isTerminated = true
	al.updatedAt = time.Now()
}

func (al *MasterAccessLog) MarkExpired() {
	al.isExpired = true
	al.updatedAt = time.Now()
}

func (al *MasterAccessLog) IncrementRefreshCount() {
	al.refreshCount++
	al.updatedAt = time.Now()
}

func (al *MasterAccessLog) IncrementRequestCount() {
	al.requestsCount++
	al.updatedAt = time.Now()
}

func (al *MasterAccessLog) IsActive() bool {
	return !al.isTerminated && !al.isExpired && time.Now().Before(al.expiresAt)
}

func (al *MasterAccessLog) CheckExpiration() {
	if time.Now().After(al.expiresAt) && !al.isExpired {
		al.MarkExpired()
	}
}

// PayloadJSON returns JSON-encoded payload
func (al *MasterAccessLog) PayloadJSON() []byte {
	if al.payload == nil {
		return []byte("{}")
	}
	data, _ := json.Marshal(al.payload)
	return data
}
