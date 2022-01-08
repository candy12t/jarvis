package twitter

import (
	"net/http"

	"github.com/candy12t/jarvis/internal/config"
	"github.com/dghubble/oauth1"
)

type Client struct {
	client *http.Client
}

func NewClient(tcfg *config.Twitter) *Client {
	cfg := oauth1.NewConfig(tcfg.ConsumerKey, tcfg.ConsumerSecret)
	token := oauth1.NewToken(tcfg.AccessToken, tcfg.AccessTokenSecret)

	return &Client{
		client: cfg.Client(oauth1.NoContext, token),
	}
}
