package main

import (
	"context"
	"io/ioutil"
	"log"
	"os"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func main() {
	if len(os.Args) < 3 {
		return
	}
	url := os.Args[1]
	out := os.Args[2]
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var buf []byte
	if err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitReady("body"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			buf , _ ,  err = page.PrintToPDF().WithPrintBackground(true).Do(ctx)
			if err != nil {
				return err
			}
			return nil
		}),
	); err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile(out, buf, 0644); err != nil {
		log.Fatal(err)
	}

	log.Println("PDF saved to example.pdf")
}