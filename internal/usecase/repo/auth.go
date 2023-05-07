package repo

import (
	"quiz-mtuci-server/pkg/logger"
	"quiz-mtuci-server/pkg/postgres"
)

type AuthRepo struct {
	*postgres.Postgres
	l *logger.Logger
}

func NewAuthRepo(pg *postgres.Postgres, l *logger.Logger) *AuthRepo {
	return &AuthRepo{pg, l}
}

func (r *AuthRepo) GetUser() error {
	return nil
}
