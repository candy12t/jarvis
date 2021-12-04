package twitter

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/mrjones/oauth"
)

const (
	requestTokenURL   = "https://api.twitter.com/oauth/request_token"
	authorizeTokenURL = "https://api.twitter.com/oauth/authorize"
	accessTokenURL    = "https://api.twitter.com/oauth/access_token"
	tweetEndpoint     = "https://api.twitter.com/2/tweets"
)

type Client struct {
	client http.Client
}

func NewClient(comsumerKey, comsumerSecret, accessToken, accessTokenSecret string) (*Client, error) {
	c := oauth.NewConsumer(
		comsumerKey,
		comsumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   requestTokenURL,
			AuthorizeTokenUrl: authorizeTokenURL,
			AccessTokenUrl:    accessTokenURL,
		},
	)

	t := &oauth.AccessToken{
		Token:  accessToken,
		Secret: accessTokenSecret,
	}

	client, err := c.MakeHttpClient(t)
	if err != nil {
		return nil, err
	}

	return &Client{client: *client}, nil
}

func (c *Client) Tweet(text string) error {
	body := map[string]string{"text": text}
	bytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	response, err := c.client.Post(tweetEndpoint, "application/json", strings.NewReader(string(bytes)))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	respBody, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	// fmt.Println(string(respBody))

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid http status")
	}

	return nil
}
