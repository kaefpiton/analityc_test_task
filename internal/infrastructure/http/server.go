package http

import (
	"analityc_test_task/internal/interfaces/httpControllers"
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
	echo                *echo.Echo
	port                string
	analitycsController httpControllers.AnalitycsController
	logger              logger.Logger
}

func NewEchoHTTPServer(
	port string,
	logger logger.Logger,
) *EchoHTTPServer {
	server := &EchoHTTPServer{
		echo:                echo.New(),
		analitycsController: httpControllers.NewAnalitycsControllerImpl(),
		port:                port,
		logger:              logger,
	}

	return server
}

func (s *EchoHTTPServer) Start() {
	s.echo.POST("/analitycs", s.handleAnalitycs)

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

func (s *EchoHTTPServer) handleAnalitycs(ctx echo.Context) error {
	return s.analitycsController.HandleAnalitics(ctx)
}
