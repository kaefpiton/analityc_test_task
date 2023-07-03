package providers

import (
	"analityc_test_task/cmd/config"
	"analityc_test_task/internal/infrastructure/http"
)

func ProvideHTTPServer(config *config.Config) http.HTTPServer {
	return http.NewEchoHTTPServer(config.HttpServer.Port)
}
