package httpControllers

import (
	"analityc_test_task/internal/entities"
	"analityc_test_task/internal/entities/api"
	"analityc_test_task/internal/usecases/interactor"
	"analityc_test_task/internal/usecases/repository"
	"analityc_test_task/pkg/logger"
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"runtime"
)

type AnalitycsController interface {
	HandleAnalitycs(c echo.Context) error
}

type analitycsController struct {
	wp         ActionsWorkerPool
	actionRepo repository.ActionsRepository
	logger     logger.Logger
}

func NewAnalitycsController(ctx context.Context, actionRepo repository.ActionsRepository, logger logger.Logger) AnalitycsController {
	actionsInteractor := interactor.NewActionsInteractor(actionRepo, logger)
	wp := NewActionsWorkerPool(ctx, runtime.NumCPU(), actionsInteractor.SendAnalityc, logger)

	return &analitycsController{
		wp:         wp,
		actionRepo: actionRepo,
		logger:     logger,
	}
}

func (c *analitycsController) HandleAnalitycs(ctx echo.Context) error {
	headers := c.getHeaders(ctx)

	body := make(map[string]interface{})
	ctx.Bind(&body)

	action := &entities.Action{
		UserId: headers["X-Tantum-Authorization"].(string),
		Data: map[string]interface{}{
			"headers": headers,
			"body":    body,
		},
	}

	c.wp.sendData(action)

	return ctx.String(http.StatusAccepted, "OK")
}

func (c *analitycsController) getHeaders(ctx echo.Context) map[string]interface{} {
	headers := make(map[string]interface{}, 0)

	headers[api.TantumAuthHeader] = ctx.Request().Header.Get(api.TantumAuthHeader)
	headers[api.ContentTypeHeader] = ctx.Request().Header.Get(api.ContentTypeHeader)
	headers[api.TantumUserAgentHeader] = ctx.Request().Header.Get(api.TantumUserAgentHeader)

	return headers
}

type ActionsWorkerPool struct {
	WorkersCount int
	receiver     chan *entities.Action
}

func NewActionsWorkerPool(ctx context.Context, workersCount int, fn handleActionFn, logger logger.Logger) ActionsWorkerPool {
	receiver := make(chan *entities.Action, workersCount)
	for i := 0; i < workersCount; i++ {
		workerID := i + 1
		go startWorker(ctx, workerID, receiver, fn, logger)
	}

	return ActionsWorkerPool{
		WorkersCount: workersCount,
		receiver:     receiver,
	}
}

type handleActionFn = func(action *entities.Action)

func startWorker(ctx context.Context, workerID int, receiver chan *entities.Action, fn handleActionFn, logger logger.Logger) {
	logger.InfoF("Start Worker [%d]", workerID)

	for {
		select {
		case data := <-receiver:
			logger.InfoF("Worker [%d] receive data = %v", workerID, data)
			fn(data)
		case <-ctx.Done():
			logger.InfoF("Shutdown worker [%d]", workerID)
			return
		default:
		}

	}
}

func (wp *ActionsWorkerPool) sendData(data *entities.Action) {
	wp.receiver <- data
}
