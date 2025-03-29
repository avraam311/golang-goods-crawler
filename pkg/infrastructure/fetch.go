package infrastructure

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/avraam311/golang-goods-crawler/internal/models/domain"
	"github.com/chromedp/chromedp"
)

func ScrapeProductDataChromeDP(url string, userAgent string, delay time.Duration) ([]domain.Product, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	req.Header.Set("User-Agent", userAgent)

	time.Sleep(delay)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute GET request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP status: %d %s", resp.StatusCode, resp.Status)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	products, err := parseHTMLContent(bytes.NewReader(bodyBytes), url)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML content: %w", err)
	}

	return products, nil
}

func parseHTMLContent(r io.Reader, url string) ([]domain.Product, error) {
	htmlContent, err := readHTMLContent(r, url)
	if err != nil {
		return nil, fmt.Errorf("failed to read HTML content: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		return nil, fmt.Errorf("failed to create goquery document: %w", err)
	}

	var products []domain.Product

	doc.Find("div.product-name").Each(func(i int, s *goquery.Selection) {
		name := s.Find("a").Text()

		priceDiv := s.Parent().Find("div.product-price")
		price := priceDiv.Find("span").Text()

		products = append(products, domain.Product{Name: strings.TrimSpace(name), Price: strings.TrimSpace(price)})
	})

	return products, nil
}

func readHTMLContent(r io.Reader, url string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	c, err := chromedp.NewContext(ctx, chromedp.WithLogf(log.Printf))
	if err != nil {
		return "", fmt.Errorf("failed to create chrome instance: %w", err)
	}
	defer chromedp.Cancel(c)

	var htmlContent string

	err = chromedp.Run(c,
		chromedp.Navigate(url),
		chromedp.WaitReady("div.product-name", chromedp.ByQuery),
		chromedp.OuterHTML("body", &htmlContent, chromedp.ByQuery),
	)
	if err != nil {
		return "", fmt.Errorf("failed to run chromedp tasks: %w", err)
	}

	return htmlContent, nil
}
