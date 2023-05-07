package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"quiz-mtuci-server/config"
	v1 "quiz-mtuci-server/internal/controller/http/v1"
	"quiz-mtuci-server/internal/usecase"
	"quiz-mtuci-server/internal/usecase/repo"
	"quiz-mtuci-server/pkg/httpserver"
	"quiz-mtuci-server/pkg/logger"
	"quiz-mtuci-server/pkg/postgres"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	pg, err := postgres.New(cfg.Postgres)

	if err != nil {
		l.Fatal().Err(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	pgA, err := postgres.New(cfg.Postgres)

	if err != nil {
		l.Fatal().Err(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pgA.Close()

	quizUseCase := usecase.New(
		l,
		repo.New(pg, l),
		repo.NewAuthRepo(pgA, l),
	)

	handler := gin.New()
	v1.NewRouter(handler, l, quizUseCase)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info().Msgf("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error().Err(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	err = httpServer.Shutdown()
	if err != nil {
		l.Error().Err(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
