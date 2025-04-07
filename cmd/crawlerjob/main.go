package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/avraam311/golang-goods-crawler/internal/config"
	"github.com/avraam311/golang-goods-crawler/internal/models/dto"
	"github.com/avraam311/golang-goods-crawler/internal/pkg/db"
	"github.com/avraam311/golang-goods-crawler/internal/pkg/parser"
	goodsrepo "github.com/avraam311/golang-goods-crawler/internal/repository/goods"
	"github.com/robfig/cron"
)

type goodsRepo interface {
	Create(ctx context.Context, good dto.Goods) error
}

type CrawlerGetJob struct {
	goodsRepo     goodsRepo
}

func main() {
	cfg := config.MustLoad()
	location, err := time.LoadLocation("Local")
	if err != nil {
		panic(err)
	}
	pool, err := db.ConnectDB(context.Background(), cfg.DBConfig)
	if err != nil {
		log.Fatalln(err)
	}
	goodsRepo := goodsrepo.New(pool)
	crawlerJob := CrawlerGetJob{
		goodsRepo: goodsRepo,
	}
	cronJob := cron.NewWithLocation(location)
	cronJob.AddJob("0 * * * *", crawlerJob)
	go func() {
		fmt.Println("Starting GetGoods() job")
		cronJob.Run()
	}()
	signals := make(chan os.Signal, 2)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGABRT, syscall.SIGTERM)
	<-signals
	fmt.Println("Cron scheduler is stopped!")
	cronJob.Stop()
}

func (job CrawlerGetJob) Run() {
	err := parser.GetGoods(job.goodsRepo, "https://www.amazon.com/s?k=gaming+mouse")

	if err != nil {
		log.Fatalln(err)
	}
}
