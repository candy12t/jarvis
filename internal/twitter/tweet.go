package twitter

import (
	"encoding/json"
	"fmt"
	"strings"
)

const tweetEndpoint = "https://api.twitter.com/2/tweets"

type PostTweetResponse struct {
	Data struct {
		ID   string `json:"id"`
		Text string `json:"text"`
	}
}

type PostTweetParams struct {
	Text string `json:"text"`
}

func (c *Client) PostTweet(params PostTweetParams) (*PostTweetResponse, error) {
	bytes, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	response, err := c.client.Post(tweetEndpoint, "application/json", strings.NewReader(string(bytes)))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	success := 200 <= response.StatusCode && response.StatusCode < 300
	if !success {
		return nil, fmt.Errorf("http error")
	}

	postTweetResponse := new(PostTweetResponse)
	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(postTweetResponse); err != nil {
		return nil, err
	}

	return postTweetResponse, nil
}
