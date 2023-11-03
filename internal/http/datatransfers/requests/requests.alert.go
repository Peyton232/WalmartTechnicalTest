package requests

import (
	"time"

	V1Domains "github.com/peyton232/walmart/internal/business/domains/v1"
)

// Create Alert Request
type CreateAlertRequest struct {
	AlertID     string    `json:"alert_id" validate:"required"`
	ServiceID   string    `json:"service_id" validate:"required"`
	ServiceName string    `json:"service_name" validate:"required"`
	Model       string    `json:"model" validate:"required"`
	AlertType   string    `json:"alert_type" validate:"required"`
	AlertTS     time.Time `json:"alert_ts" validate:"required"`
	Severity    string    `json:"severity" validate:"required"`
	TeamSlack   string    `json:"team_slack" validate:"required"`
}

// Mapping Create Alert Request to Domain
func (alert CreateAlertRequest) ToV1Domain() *V1Domains.AlertDomain {
	return &V1Domains.AlertDomain{
		AlertID:     alert.AlertID,
		ServiceID:   alert.ServiceID,
		ServiceName: alert.ServiceName,
		Model:       alert.Model,
		AlertType:   alert.AlertType,
		AlertTS:     alert.AlertTS,
		Severity:    alert.Severity,
		TeamSlack:   alert.TeamSlack,
		CreatedAt:   time.Now(),
	}
}

// Read Alert Request
type ReadAlertRequest struct {
	ServiceID string `json:"service_id" validate:"required"`
	StartTS   int64  `json:"start_ts" validate:"required"`
	EndTS     int64  `json:"end_ts" validate:"required"`
}
