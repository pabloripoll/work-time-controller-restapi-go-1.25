package dto

import "time"

type MasterAccessLogDTO struct {
	ID            int64                  `json:"id"`
	UserID        string                 `json:"user_id"`
	IsTerminated  bool                   `json:"is_terminated"`
	IsExpired     bool                   `json:"is_expired"`
	ExpiresAt     time.Time              `json:"expires_at"`
	RefreshCount  int                    `json:"refresh_count"`
	CreatedAt     time.Time              `json:"created_at"`
	UpdatedAt     time.Time              `json:"updated_at"`
	IPAddress     string                 `json:"ip_address"`
	UserAgent     string                 `json:"user_agent"`
	RequestsCount int                    `json:"requests_count"`
	Payload       map[string]interface{} `json:"payload,omitempty"`
}
