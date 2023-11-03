package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	V1Usecase "github.com/peyton232/walmart/internal/business/usecases/v1"
	V1PostgresRepository "github.com/peyton232/walmart/internal/datasources/repositories/postgres/v1"
	V1Handler "github.com/peyton232/walmart/internal/http/handlers/v1"
)

type alertRoutes struct {
	V1Handler V1Handler.AlertHandler
	router    *gin.RouterGroup
	db        *sqlx.DB
}

func NewAlertRoute(router *gin.RouterGroup, db *sqlx.DB) *alertRoutes {
	V1AlertRepository := V1PostgresRepository.NewUserRepository(db)
	V1AlertUsecase := V1Usecase.NewAlertUsecase(V1AlertRepository)
	V1UserHandler := V1Handler.NewAlertHandler(V1AlertUsecase)

	return &alertRoutes{V1Handler: V1UserHandler, router: router, db: db}
}

func (r *alertRoutes) Routes() {
	// Routes V1
	V1Route := r.router.Group("/v1")
	{
		// alert
		V1AuhtRoute := V1Route.Group("/alerts")
		V1AuhtRoute.POST("", r.V1Handler.Regis)
		V1AuhtRoute.GET("", r.V1Handler.GetAlertsByService)
	}

}
