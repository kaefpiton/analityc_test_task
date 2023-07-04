package providers

import (
	"analityc_test_task/cmd/config"
	"analityc_test_task/internal/infrastructure/http"
	"analityc_test_task/pkg/db/postgres"
	"analityc_test_task/pkg/logger"
	"analityc_test_task/pkg/logger/zerolog"
	"os"
)

func ProvideHTTPServer(config *config.Config, logger logger.Logger) http.HTTPServer {
	return http.NewEchoHTTPServer(config.HttpServer.Port, logger)
}

func ProvideConsoleLogger(cnf *config.Config) (logger.Logger, error) {
	return zerolog.NewZeroLog(os.Stderr, cnf.Logger.Lvl)
}

// todo вернуть closer func
func ProvideDB(cnf *config.Config, log logger.Logger) (*postgres.DB, error) {
	repo, err := postgres.NewDBConnection(cnf, log)
	if err != nil {
		return nil, err
	}

	return repo, nil
}
