package http

import (
	"analityc_test_task/internal/entities/api"
	"analityc_test_task/internal/interfaces/httpControllers"
	"analityc_test_task/pkg/logger"
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type HTTPServer interface {
	Start()
	Stop(ctx context.Context)
}

type EchoHTTPServer struct {
	echo                *echo.Echo
	port                string
	analitycsController httpControllers.AnalitycsController
	logger              logger.Logger
}

func NewEchoHTTPServer(
	port string,
	analitycsController httpControllers.AnalitycsController,
	logger logger.Logger,
) *EchoHTTPServer {
	server := &EchoHTTPServer{
		echo: echo.New(),
		//todo вынести в провайдер
		analitycsController: analitycsController,
		port:                port,
		logger:              logger,
	}

	return server
}

func (s *EchoHTTPServer) Start() {
	s.echo.Use(ValidateHeaders)
	s.echo.POST("/analitycs", s.handleAnalitycs)

	func() {
		port := fmt.Sprintf(":%v", s.port)
		if err := s.echo.Start(port); err != nil {
			s.logger.Error("Echo error:", err)
		}
	}()
}

func (s *EchoHTTPServer) Stop(ctx context.Context) {
	err := s.echo.Shutdown(ctx)
	if err != nil {
		s.logger.Error("Echo error:", err)
	}
}

func ValidateHeaders(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Header.Get(api.ContentTypeHeader) == "" {
			return c.String(http.StatusBadRequest, api.InvalidHeadersError)
		}

		if c.Request().Header.Get(api.TantumAuthHeader) == "" {
			return c.String(http.StatusBadRequest, api.InvalidHeadersError)
		}

		if c.Request().Header.Get(api.TantumUserAgentHeader) == "" {
			return c.String(http.StatusBadRequest, api.InvalidHeadersError)
		}

		return next(c)
	}
}

func (s *EchoHTTPServer) handleAnalitycs(ctx echo.Context) error {
	return s.analitycsController.HandleAnalitycs(ctx)
}
