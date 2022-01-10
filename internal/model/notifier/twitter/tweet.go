package twitter

import (
	"context"
	"fmt"
	"net/http"
)

type PostTweetResponse struct {
	Data struct {
		ID   string `json:"id"`
		Text string `json:"text"`
	}
}

type PostTweetParams struct {
	Text string `json:"text"`
}

func (c *Client) PostTweet(ctx context.Context, params PostTweetParams) (*PostTweetResponse, *http.Response, error) {
	u := fmt.Sprintf("tweets")
	req, err := c.NewRequest("POST", u, params)
	if err != nil {
		return nil, nil, err
	}

	t := new(PostTweetResponse)
	resp, err := c.Do(ctx, req, t)
	if err != nil {
		return nil, nil, err
	}

	return t, resp, nil
}
