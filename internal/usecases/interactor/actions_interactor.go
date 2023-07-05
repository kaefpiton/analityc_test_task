package interactor

import (
	"analityc_test_task/internal/entities"
	"analityc_test_task/internal/usecases/repository"
	"analityc_test_task/pkg/db/postgres"
	"analityc_test_task/pkg/logger"
	"encoding/json"
)

type ActionsInteractor interface {
	SendAnalityc(action *entities.Action)
}

type actionsInteractor struct {
	db          *postgres.DB
	actionsRepo repository.ActionsRepository
	logger      logger.Logger
}

func NewActionsInteractor(actionRepo repository.ActionsRepository, logger logger.Logger) ActionsInteractor {
	return &actionsInteractor{
		actionsRepo: actionRepo,
		logger:      logger,
	}
}

func (i *actionsInteractor) SendAnalityc(action *entities.Action) {
	b, err := json.Marshal(action.Data)
	if err != nil {
		i.logger.Error("Actions Interactor error:", err)
	}

	i.actionsRepo.Create(action.UserId, b)
}
