package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/avraam311/golang-goods-crawler/internal/api"
	"github.com/avraam311/golang-goods-crawler/internal/config"
	"github.com/avraam311/golang-goods-crawler/internal/pkg/db"
	"github.com/avraam311/golang-goods-crawler/internal/pkg/logger"
	goodsrepo "github.com/avraam311/golang-goods-crawler/internal/repository/goods"
	"github.com/avraam311/golang-goods-crawler/internal/service/goods"
)

func main() {
	cfg := config.MustLoad()
	pool, err := db.ConnectDB(context.Background(), cfg.DBConfig)
	if err != nil {
		log.Fatalln(err)
	}
	myLogger, err := logger.New(cfg.LogFilePath)
	if err != nil {
		log.Fatalln(err)
	}
	goodsRepo := goodsrepo.New(pool)
	goodsService := goods.New(goodsRepo)
	if err != nil {
		log.Fatalln(err)
	}
	r, err := api.New(
		myLogger,
		goodsService,
	)
	if err != nil {
		log.Fatalln(err)
	}

	srv := &http.Server{
		Addr:    cfg.Address,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			fmt.Println("Failed to start http server!")
		}
		fmt.Println("Listening on localhost:3000")
	}()
	signals := make(chan os.Signal, 2)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGABRT, syscall.SIGTERM)
	<-signals
	fmt.Println("Server stoped!")
	srv.Shutdown(context.Background())
}
