package v1

import (
	"quiz-mtuci-server/internal/usecase"
	"quiz-mtuci-server/pkg/logger"
	"quiz-mtuci-server/pkg/metrics"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewRouter(handler *gin.Engine, l logger.Interface, t usecase.UseCase) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Prometheus metrics
	handler.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Routers
	h := handler.Group("/v1")
	{
		newQuizRoutes(h, t, l)
	}

	// register metrics
	if err := metrics.RegisterMetrics(); err != nil {
		l.Error().Err(err)

		return
	}
}
