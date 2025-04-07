package parser

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/avraam311/golang-goods-crawler/internal/models/dto"
	"github.com/chromedp/chromedp"
	"github.com/jackc/pgx/v5"
)

type GoodsInterface interface {
	Create(ctx context.Context, good dto.Goods) error
}

func GetGoods(goodsRepo GoodsInterface, url string) error {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 600*time.Second)
	defer cancel()

	var productTitles []string

	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible(`h2.a-size-medium.a-spacing-none.a-color-base.a-text-normal`, chromedp.ByQueryAll),
		chromedp.Evaluate(`Array.from(document.querySelectorAll("h2.a-size-medium.a-spacing-none.a-color-base.a-text-normal")).map((el) => el.innerText);`, &productTitles),
	)
	if err != nil {
		return err
	}

	for _, title := range productTitles {
		good := dto.Goods{Name: title}
		err = goodsRepo.Create(context.Background(), good)
		if !errors.Is(err, pgx.ErrNoRows) {
			fmt.Printf("parser/goods.go.GetGoods: %s\n", err)
			return err
		}
	}

	return nil
}
