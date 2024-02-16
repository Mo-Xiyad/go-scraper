package main

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	log.Println("Starting ChromeDP...")

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
	)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()


	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 30*time.Second) 
	defer cancel()

	var providerName, providerAddress string

	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://www.foodora.se/restaurant/s7gw/thai-n-sushi-for-you`),
		chromedp.WaitVisible(`[data-testid="vendor-info-more-info-btn"]`),
		chromedp.Click(`[data-testid="vendor-info-more-info-btn"]`, chromedp.NodeVisible),
		chromedp.Text(`[data-testid="provider-legal-name"]`, &providerName),
		chromedp.Text(`[data-testid="provider-info-address"]`, &providerAddress),
	)
	if err != nil {
		log.Fatalf("Failed to run ChromeDP tasks: %v", err)
	}

	log.Printf("provider name: %s, provider address: %s", providerName, providerAddress)

}
