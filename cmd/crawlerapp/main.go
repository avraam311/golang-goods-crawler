package main

import (
	"fmt"
	"log"
	"time"

	"github.com/avraam311/golang-goods-crawler/pkg/infrastructure"
)

func main() {
	targetURL := "https://www.okeydostavka.ru/msk"
	userAgent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36"
	delay := 2 * time.Second

	products, err := infrastructure.ScrapeProductDataChromeDP(targetURL, userAgent, delay)
	if err != nil {
		log.Fatalf("Failed to scrape and parse data: %v", err)
	}

	if len(products) == 0 {
		fmt.Println("No products found.")
		return
	}

	fmt.Println("Found products:")
	for i, product := range products {
		fmt.Printf("%d. Name: %s, Price: %s\n", i+1, product.Name, product.Price)
	}

	fmt.Println("Data scraping completed.")
}
