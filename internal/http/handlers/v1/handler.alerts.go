package v1

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	V1Domains "github.com/peyton232/walmart/internal/business/domains/v1"
	"github.com/peyton232/walmart/internal/http/datatransfers/requests"
	"github.com/peyton232/walmart/internal/http/datatransfers/responses"
	"github.com/peyton232/walmart/pkg/validators"
)

type AlertHandler struct {
	usecase V1Domains.AlertUsecase
}

func NewAlertHandler(usecase V1Domains.AlertUsecase) AlertHandler {
	return AlertHandler{
		usecase: usecase,
	}
}

// TODO rename
func (alertH AlertHandler) Regis(ctx *gin.Context) {
	var AlertRegisRequest requests.CreateAlertRequest
	if err := ctx.ShouldBindJSON(&AlertRegisRequest); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := validators.ValidatePayloads(AlertRegisRequest); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	alertDomain := AlertRegisRequest.ToV1Domain()
	statusCode, err := alertH.usecase.CreateAlert(ctx.Request.Context(), alertDomain)
	fmt.Println(alertDomain, statusCode, err)
	if err != nil {
		NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	NewSuccessResponse(ctx, statusCode, "created alert success", map[string]interface{}{
		"alert": alertDomain,
	})
}

func (a AlertHandler) GetAlertsByService(ctx *gin.Context) {

	var AlertRequest requests.ReadAlertRequest
	if err := ctx.ShouldBindJSON(&AlertRequest); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctxx := ctx.Request.Context()
	alertDom, statusCode, err := a.usecase.ReadAlerts(ctxx, AlertRequest.ServiceID, time.Unix(AlertRequest.StartTS, 0), time.Unix(AlertRequest.EndTS, 0))
	if err != nil {
		NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	alertResponse := responses.ToResponseList(alertDom)

	NewSuccessResponse(ctx, statusCode, "alert data fetched successfully", map[string]interface{}{
		"alert": alertResponse,
	})

}
