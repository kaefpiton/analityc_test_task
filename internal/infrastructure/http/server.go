package http

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
)

type HTTPServer interface {
	Start()
	Stop()
}

type EchoHTTPServer struct {
	echo *echo.Echo
	//todo прокинуть логгер

	port string
}

func NewEchoHTTPServer(
	port string,
) *EchoHTTPServer {
	server := &EchoHTTPServer{
		echo: echo.New(),
		port: port,
	}

	return server
}

func (s *EchoHTTPServer) Start() {
	func() {
		port := fmt.Sprintf(":%v", s.port)
		if err := s.echo.Start(port); err != nil {
			//todo запровайдить лог, вывести ошибку
		}
	}()
}

// todo для шатдауна
func (s *EchoHTTPServer) Stop() {
	err := s.echo.Shutdown(context.Background())
	if err != nil {
		//todo запровайдить лог, вывести ошибку
	}
}
