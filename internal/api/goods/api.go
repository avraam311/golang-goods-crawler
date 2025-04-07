package goods

import (
	"context"

	"github.com/avraam311/golang-goods-crawler/internal/models/domain"
)

type goodsService interface {
	GetGoods(ctx context.Context) ([]domain.Goods, error)
}

type logger interface {
	Info(text ...any)
	Warn(text ...any)
	Err(text ...any)
}

type API struct {
	logger  logger
	service goodsService
}

func NewAPI(logger logger, service goodsService) *API {
	return &API{
		logger:  logger,
		service: service,
	}
}
