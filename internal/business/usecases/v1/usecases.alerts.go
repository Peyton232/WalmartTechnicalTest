package v1

import (
	"context"
	"net/http"
	"time"

	V1Domains "github.com/peyton232/walmart/internal/business/domains/v1"
	"github.com/peyton232/walmart/internal/constants"
)

type alertUsecase struct {
	repo V1Domains.AlertRepository
}

func NewAlertUsecase(repo V1Domains.AlertRepository) V1Domains.AlertUsecase {
	return &alertUsecase{
		repo: repo,
	}
}

func (alertUC *alertUsecase) CreateAlert(ctx context.Context, inDom *V1Domains.AlertDomain) (statusCode int, err error) {
	inDom.CreatedAt = time.Now().In(constants.GMT7)

	err = alertUC.repo.CreateAlert(ctx, inDom)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}

func (alertUC *alertUsecase) ReadAlerts(ctx context.Context, serviceID string, startTS time.Time, endTS time.Time) (outDom []V1Domains.AlertDomain, statusCode int, err error) {
	outDom, err = alertUC.repo.GetAlertsByTimeStamp(ctx, serviceID, startTS, endTS)

	if err != nil {
		return outDom, http.StatusInternalServerError, err
	}

	return outDom, http.StatusOK, nil
}
