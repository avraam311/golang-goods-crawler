package main

import (
	"fmt"
	"log"

	"github.com/avraam311/golang-goods-crawler/pkg/infrastructure"
)

func main() {
    url := "https://example.com" // Замените на нужный URL
    productName, err := FetchProductName(url)
    if err != nil {
        log.Fatalf("Error fetching product name: %v", err)
    }

    fmt.Printf("Product Name: %s\n", productName)
}
