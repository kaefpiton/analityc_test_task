package repository

import (
	"analityc_test_task/internal/usecases/repository"
	"analityc_test_task/pkg/db/postgres"
	"analityc_test_task/pkg/logger"
	"sync"
)

type actionsRepository struct {
	db     *postgres.DB
	mu     sync.Mutex
	logger logger.Logger
}

func NewActionRepository(db *postgres.DB, logger logger.Logger) repository.ActionsRepository {
	return &actionsRepository{
		db:     db,
		logger: logger,
	}
}

func (r *actionsRepository) Create(userId string, data []byte) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	query := "INSERT INTO actions (user_id, data) values ($1, $2)"
	_, err := r.db.Exec(query, userId, data)
	if err != nil {
		r.logger.ErrorF("repo error:%w", err)
	}

	return err
}
