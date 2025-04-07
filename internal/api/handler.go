package api

import (
	"context"

	"github.com/avraam311/golang-goods-crawler/internal/api/goods"
	"github.com/avraam311/golang-goods-crawler/internal/models/domain"
	"github.com/gin-gonic/gin"
)

type goodsService interface {
	GetGoods(ctx context.Context) ([]domain.Goods, error)
}

type logger interface {
	Info(text ...any)
	Warn(text ...any)
	Err(text ...any)
}

func New(logger logger, goodsService goodsService) (*gin.Engine, error) {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(CORSMiddleware())

	goodsAPI := goods.NewAPI(logger, goodsService)
	api := r.Group("/")
	api.GET("/goods", goodsAPI.GetGoods)
	return r, nil
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)

			return
		}

		c.Next()
	}
}
