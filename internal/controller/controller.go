package controller

import (
	"context"
	"fmt"

	"github.com/candy12t/n3m/internal/scraping"
	"github.com/candy12t/n3m/internal/twitter"
	"github.com/chromedp/chromedp"
)

func TweetNewMedia(client twitter.Client) error {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	hrefs, err := scraping.GetNetflixNode(ctx)
	if err != nil {
		return err
	}

	titles, err := scraping.GetTitles(ctx, hrefs)
	if err != nil {
		return err
	}

	message := "新着映画"
	for _, title := range titles {
		message += fmt.Sprintf("\n%s", title)
	}
	if err := client.Tweet(message); err != nil {
		return err
	}
	return nil
}
