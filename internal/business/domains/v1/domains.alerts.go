package v1

import (
	"context"
	"time"
)

type AlertDomain struct {
	AlertID     string
	ServiceID   string
	ServiceName string
	Model       string
	AlertType   string
	AlertTS     time.Time
	Severity    string
	TeamSlack   string
	CreatedAt   time.Time
}

type AlertUsecase interface {
	CreateAlert(ctx context.Context, inDom *AlertDomain) (statusCode int, err error)
	ReadAlerts(ctx context.Context, serviceID string, startTS time.Time, endTS time.Time) (outDom []AlertDomain, statusCode int, err error)
}

type AlertRepository interface {
	Store(ctx context.Context, inDom *AlertDomain) (err error)
	GetAlertsByTimeStamp(ctx context.Context, serviceID string, startTS time.Time, endTS time.Time) (outDomain []AlertDomain, err error)
}
