package http

import (
	"analityc_test_task/pkg/logger"
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
)

type HTTPServer interface {
	Start()
	Stop()
}

type EchoHTTPServer struct {
	echo   *echo.Echo
	logger logger.Logger
	port   string
}

func NewEchoHTTPServer(
	port string,
	logger logger.Logger,
) *EchoHTTPServer {
	server := &EchoHTTPServer{
		echo:   echo.New(),
		port:   port,
		logger: logger,
	}

	return server
}

func (s *EchoHTTPServer) Start() {
	func() {
		port := fmt.Sprintf(":%v", s.port)
		if err := s.echo.Start(port); err != nil {
			s.logger.Error("Echo error:", err)
		}
	}()
}

// todo для шатдауна
func (s *EchoHTTPServer) Stop() {
	err := s.echo.Shutdown(context.Background())
	if err != nil {
		s.logger.Error("Echo error:", err)
	}
}
