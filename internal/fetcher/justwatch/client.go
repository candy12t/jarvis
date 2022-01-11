package justwatch

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

const (
	defaultBaseURL = "https://apis.justwatch.com"
)

type Client struct {
	client  *http.Client
	BaseURL *url.URL
	locale  string
}

func NewClient(locale string) *Client {
	baseURL, _ := url.Parse(defaultBaseURL)
	return &Client{
		client:  http.DefaultClient,
		BaseURL: baseURL,
		locale:  locale,
	}
}

func (c *Client) NewRequest(method, urlStr string, body io.Reader) (*http.Request, error) {
	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
		return nil, err
	}
	return resp, nil
}

func addOptions(s string, opts interface{}) (string, error) {
	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	q, err := query.Values(opts)
	if err != nil {
		return s, err
	}

	u.RawQuery = q.Encode()
	return u.String(), nil
}
