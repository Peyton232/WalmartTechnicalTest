package responses

import (
	"time"

	V1Domains "github.com/peyton232/walmart/internal/business/domains/v1"
)

type AlertResponse struct {
	AlertID   string    `json:"alert_id" validate:"required"`
	Model     string    `json:"model" validate:"required"`
	AlertType string    `json:"alert_type" validate:"required"`
	AlertTS   time.Time `json:"alert_ts" validate:"required"`
	Severity  string    `json:"severity" validate:"required"`
	TeamSlack string    `json:"team_slack" validate:"required"`
}

func (u *AlertResponse) ToV1Domain() V1Domains.AlertDomain {
	return V1Domains.AlertDomain{
		AlertID:   u.AlertID,
		Model:     u.Model,
		AlertType: u.AlertID,
		AlertTS:   u.AlertTS,
		Severity:  u.Severity,
		TeamSlack: u.TeamSlack,
	}
}

func FromV1Domain(u V1Domains.AlertDomain) AlertResponse {
	return AlertResponse{
		AlertID:   u.AlertID,
		Model:     u.Model,
		AlertType: u.AlertType,
		AlertTS:   u.AlertTS,
	}
}

func ToResponseList(domains []V1Domains.AlertDomain) []AlertResponse {
	var result []AlertResponse

	for _, val := range domains {
		result = append(result, FromV1Domain(val))
	}

	return result
}
