package infrastructure

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

// FetchProductName получает название продукта с указанного URL.
func FetchProductName(url string) (string, error) {
    // Создаем контекст с таймаутом
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    // Создаем новый контекст для chromedp
    c, err := chromedp.NewContext(ctx)
    if err != nil {
        return "", fmt.Errorf("failed to create context: %w", err)
    }

    var productName string

    // Выполняем задачи chromedp
    err = chromedp.Run(c,
        chromedp.Navigate(url),
        chromedp.WaitVisible("div.product-name", chromedp.ByQuery), // Ждем, пока элемент станет видимым
        chromedp.Text("div.product-name", &productName, chromedp.ByQuery), // Получаем текст из элемента
    )
    if err != nil {
        return "", fmt.Errorf("failed to run tasks: %w", err)
    }

    return productName, nil
}

func main() {
    url := "https://example.com" // Замените на нужный URL
    productName, err := FetchProductName(url)
    if err != nil {
        log.Fatalf("Error fetching product name: %v", err)
    }

    fmt.Printf("Product Name: %s\n", productName)
}
