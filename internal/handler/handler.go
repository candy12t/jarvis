package handler

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/candy12t/jarvis/internal/config"
	"github.com/candy12t/jarvis/internal/justwatch"
	"github.com/candy12t/jarvis/internal/twitter"
)

type Contents []Content

type Content struct {
	Title      string
	ContentURL string
}

const dateFormat = "2006-01-02"

func Apply() error {
	// timeNow := time.Now().Format(dateFormat)
	timeNow := "2022-01-08"

	contents, err := getContents(timeNow)
	if err != nil {
		return err
	}

	return noticeContens(timeNow, contents)
}

func getContents(date string) (Contents, error) {
	notes := make(Contents, 0)
	client := justwatch.NewClient(config.JustwatchCfg().Country)

	for page := 1; ; page++ {
		opts := &justwatch.ContentBodyOptions{
			Providers:    config.JustwatchCfg().Services,
			Date:         date,
			ContentTypes: []string{"movie"},
			Page:         page,
			PageSize:     5,
		}

		ctx := context.Background()
		contents, err := client.ListNewContens(ctx, opts)
		if err != nil {
			return nil, err
		}

		if contents.TotalPages == 0 {
			client := twitter.NewClient(config.TwitterCfg())
			params := twitter.PostTweetParams{
				Text: fmt.Sprintf("%s 新着なし", date),
			}
			resp, err := client.PostTweet(params)
			if err != nil {
				return nil, err
			}
			log.Printf("ID: %v, Text: %v\n", resp.Data.ID, strings.Replace(resp.Data.Text, "\n", " ", 3))
			break
		}

		for _, content := range contents.Items {
			for _, offer := range content.Offers {
				note := Content{
					Title:      content.Title,
					ContentURL: offer.Urls.StandardWeb,
				}
				notes = append(notes, note)
			}
		}

		if contents.Page == contents.TotalPages {
			break
		}
	}

	return notes, nil
}

func noticeContens(date string, contents Contents) error {
	client := twitter.NewClient(config.TwitterCfg())
	for _, content := range contents {
		text := fmt.Sprintf("%s\n%s\n%s", fmt.Sprintf("%s", date), content.Title, content.ContentURL)
		params := twitter.PostTweetParams{
			Text: text,
		}
		resp, err := client.PostTweet(params)
		if err != nil {
			return err
		}
		log.Printf("ID: %v, Text: %v\n", resp.Data.ID, strings.Replace(resp.Data.Text, "\n", " ", 3))
	}
	return nil
}
