package goods

import (
	"context"

	"github.com/avraam311/golang-goods-crawler/internal/models/domain"
	"github.com/avraam311/golang-goods-crawler/internal/models/dto"
	"github.com/avraam311/golang-goods-crawler/internal/service"
)

type goodsRepo interface {
	Create(ctx context.Context, good dto.Goods) error
	GetAll(ctx context.Context) ([]domain.Goods, error)
}

type Service struct {
	goodsRepo goodsRepo
}

func New(goodsRepo goodsRepo) *Service {
	return &Service{
		goodsRepo: goodsRepo,
	}
}

func (s *Service) GetGoods(ctx context.Context) ([]domain.Goods, error) {
	goods, err := s.goodsRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	if goods == nil {
		return nil, service.ErrGoodsListIsEmpty
	}
	return goods, nil
}
