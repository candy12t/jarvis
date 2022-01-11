package justwatch

import (
	"context"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/candy12t/jarvis/internal/config"
	"github.com/candy12t/jarvis/internal/fetcher"
	"github.com/candy12t/jarvis/internal/model"
)

var _ fetcher.Fetcher = &Client{}

func (c *Client) FetchContents(date string) (model.Contents, error) {
	notes := make(model.Contents, 0)
	for page := 1; ; page++ {
		opts := &ContentBodyOptions{
			Providers:    config.JustwatchProviders(),
			Date:         date,
			ContentTypes: []string{"movie"},
			Page:         page,
			PageSize:     5,
		}

		ctx := context.Background()
		contents, err := c.ListNewContens(ctx, opts)
		if err != nil {
			return nil, err
		}

		if contents.TotalPages == 0 {
			return nil, nil
		}

		for _, content := range contents.Items {
			for _, offer := range content.Offers {
				contentURL := normalizationURL(offer.Urls)
				note := model.Content{
					Title:      content.Title,
					ContentURL: contentURL,
				}
				log.Printf("[justwatch]: %+v\n", note)
				notes = append(notes, note)
			}
		}

		if contents.Page == contents.TotalPages {
			break
		}
	}

	return notes, nil
}

func normalizationURL(urls Urls) string {
	if strings.Contains(urls.StandardWeb, "netflix") {
		return urls.StandardWeb
	}
	if strings.Contains(urls.StandardWeb, "amazon") {
		return getLocation(urls.StandardWeb)
	}
	if strings.Contains(urls.StandardWeb, "disneyplus") {
		return urls.DeeplinkWeb
	}
	return urls.StandardWeb
}

func getLocation(_url string) string {
	userAgent := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7)"
	req, err := http.NewRequest("GET", _url, nil)
	if err != nil {
		return _url
	}
	req.Header.Set("user-agent", userAgent)

	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		return _url
	}
	defer func() {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}()

	if 300 <= resp.StatusCode && resp.StatusCode < 400 {
		location, err := resp.Location()
		if err != nil {
			return _url
		}
		return location.String()
	}
	return _url
}
