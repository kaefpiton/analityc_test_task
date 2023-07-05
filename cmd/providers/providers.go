package providers

import (
	"analityc_test_task/configs"
	"analityc_test_task/internal/infrastructure/http"
	repository2 "analityc_test_task/internal/infrastructure/repository"
	"analityc_test_task/internal/interfaces/httpControllers"
	"analityc_test_task/internal/usecases/repository"
	"analityc_test_task/pkg/db/postgres"
	"analityc_test_task/pkg/logger"
	"analityc_test_task/pkg/logger/zerolog"
	"context"
	"os"
)

func ProvideHTTPServer(config *configs.Config, analitycsController httpControllers.AnalitycsController, logger logger.Logger) http.HTTPServer {
	return http.NewEchoHTTPServer(config.HttpServer.Port, config.HttpServer.MetricsPort, analitycsController, logger)
}

func ProvideConsoleLogger(cnf *configs.Config) (logger.Logger, error) {
	return zerolog.NewZeroLog(os.Stderr, cnf.Logger.Lvl)
}

func ProvideDB(cnf *configs.Config, logger logger.Logger) (*postgres.DB, func(), error) {
	repo, err := postgres.NewDBConnection(cnf, logger)
	if err != nil {
		return nil, nil, err
	}
	closer := func() {
		repo.Close()
	}

	return repo, closer, nil
}

func ProvideActionRepository(db *postgres.DB, logger logger.Logger) repository.ActionsRepository {
	return repository2.NewActionRepository(db, logger)
}

func ProvideAnalitycsController(ctx context.Context, actionRepo repository.ActionsRepository, logger logger.Logger) httpControllers.AnalitycsController {
	return httpControllers.NewAnalitycsController(ctx, actionRepo, logger)
}
