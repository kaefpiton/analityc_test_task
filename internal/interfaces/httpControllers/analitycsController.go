package httpControllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type AnalitycsController interface {
	HandleAnalitics(c echo.Context) error
}

type analitycsControllerImpl struct {
}

func NewAnalitycsControllerImpl() *analitycsControllerImpl {
	return &analitycsControllerImpl{}
}

func (c *analitycsControllerImpl) HandleAnalitics(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "well done")
}
