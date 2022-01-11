package twitter

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/dghubble/oauth1"
)

const (
	defaultBaseURL = "https://api.twitter.com/2/"
	userAgent      = "jarvis"
)

type Client struct {
	client    *http.Client
	baseURL   *url.URL
	UserAgent string
}

func NewAuthenticator(consumerKey, consumerSecret, accessToken, accessTokenSecret string) *http.Client {
	cfg := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessTokenSecret)
	return cfg.Client(oauth1.NoContext, token)
}

func NewClient(httpClient *http.Client) *Client {
	baseURL, _ := url.Parse(defaultBaseURL)
	return &Client{
		client:    httpClient,
		baseURL:   baseURL,
		UserAgent: userAgent,
	}
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	u, err := c.baseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	return req, nil
}

func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	success := 200 <= resp.StatusCode && resp.StatusCode < 300
	if !success {
		return resp, checkResponse(resp)
	}

	switch v := v.(type) {
	case nil:
		_, err = io.Copy(ioutil.Discard, resp.Body)
	case io.Writer:
		_, err = io.Copy(v, resp.Body)
	default:
		err = json.NewDecoder(resp.Body).Decode(v)
	}
	return resp, err
}
