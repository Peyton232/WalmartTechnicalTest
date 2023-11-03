package records

import (
	"time"
)

type Alerts struct {
	AlertID     string    `db:"alert_id"`
	ServiceID   string    `db:"service_id"`
	ServiceName string    `db:"service_name"`
	Model       string    `db:"model"`
	AlertType   string    `db:"alert_type"`
	AlertTS     time.Time `db:"alert_ts"`
	Severity    string    `db:"severity"`
	TeamSlack   string    `db:"team_slack"`
	CreatedAt   time.Time `db:"created_at"`
}
