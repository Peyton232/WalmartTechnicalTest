package v1

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	V1Domains "github.com/peyton232/walmart/internal/business/domains/v1"
	"github.com/peyton232/walmart/internal/datasources/records"
)

type postgreAlertRepository struct {
	conn *sqlx.DB
}

func NewUserRepository(conn *sqlx.DB) V1Domains.AlertRepository {
	return &postgreAlertRepository{
		conn: conn,
	}
}

func (r *postgreAlertRepository) Store(ctx context.Context, inDom *V1Domains.AlertDomain) (err error) {
	alertRecord := records.FromUsersV1Domain(inDom)

	_, err = r.conn.NamedQueryContext(ctx, `INSERT INTO alerts(alert_id, service_id, service_name, model, alert_type, alert_ts, severity, team_slack, created_at) VALUES (:alert_id, :service_id, :service_name, :model, :alert_type, :alert_ts, :severity, :team_slack, :created_at)`, alertRecord)
	if err != nil {
		return err
	}

	return nil
}

func (r *postgreAlertRepository) GetAlertsByTimeStamp(ctx context.Context, serviceID string, startTS time.Time, endTS time.Time) (outDomain []V1Domains.AlertDomain, err error) {
	var alertRecords *[]records.Alerts

	err = r.conn.GetContext(ctx, &alertRecords, `SELECT * FROM alerts WHERE service_id = $1 AND alert_ts >= $2 AND alert_ts <= $3`, serviceID, startTS, endTS)
	if err != nil {
		return []V1Domains.AlertDomain{}, err
	}

	return records.ToArrayOfUsersV1Domain(alertRecords), nil
}
