package twitter

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/candy12t/jarvis/internal/model"
	"github.com/candy12t/jarvis/internal/notifier"
)

var _ notifier.Notifier = &Client{}

func (c *Client) NotifyContents(date string, contents model.Contents) error {
	for _, content := range contents {
		text := fmt.Sprintf("%s\n%s\n%s", fmt.Sprintf("%s", date), content.Title, content.ContentURL)

		resp, err := c.notify(text)
		if err != nil {
			return err
		}
		log.Printf("[twitter]: ID: %v, Text: %v\n", resp.Data.ID, strings.Replace(resp.Data.Text, "\n", " ", 3))
	}
	return nil
}

func (c *Client) NotifyNoContents(date string) error {
	text := fmt.Sprintf("%s 新着なし", date)

	resp, err := c.notify(text)
	if err != nil {
		return err
	}
	log.Printf("[twitter]: ID: %v, Text: %v\n", resp.Data.ID, strings.Replace(resp.Data.Text, "\n", " ", 3))
	return nil
}

func (c *Client) notify(text string) (*PostTweetResponse, error) {
	ctx := context.Background()
	params := PostTweetParams{Text: text}
	postTweetResp, _, err := c.PostTweet(ctx, params)
	return postTweetResp, err
}
