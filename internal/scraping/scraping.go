package scraping

import (
	"context"
	"fmt"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

const justwatchEndpoint = "https://www.justwatch.com/jp/%E5%8B%95%E7%94%BB%E9%85%8D%E4%BF%A1%E3%82%B5%E3%83%BC%E3%83%93%E3%82%B9/netflix/%E6%96%B0%E4%BD%9C"
const dateFormt = "2006-01-02"
const justwatchDomain = "https://www.justwatch.com"

func GetNetflixNode(ctx context.Context) ([]*cdp.Node, error) {
	var nodes []*cdp.Node
	err := chromedp.Run(
		ctx,
		chromedp.Navigate(justwatchEndpoint),
		chromedp.Nodes(todayNetflixTag(), &nodes),
	)
	if err != nil {
		return nil, err
	}
	return getAtag(ctx, nodes)
}

func getAtag(ctx context.Context, nodes []*cdp.Node) ([]*cdp.Node, error) {
	var childrenNode []*cdp.Node
	err := chromedp.Run(
		ctx,
		chromedp.Nodes("a", &childrenNode, chromedp.ByQueryAll, chromedp.FromNode(nodes[0])),
	)
	if err != nil {
		return nil, err
	}
	return childrenNode, nil
}

func GetTitles(ctx context.Context, childrenNode []*cdp.Node) ([]string, error) {
	titles := make([]string, 0)
	for _, child := range childrenNode {
		absLink := absPath(child)
		title, err := getTitle(absLink)
		if err != nil {
			return nil, err
		}
		titles = append(titles, title)
	}
	return titles, nil
}

func getTitle(absLink string) (string, error) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var title string
	err := chromedp.Run(
		ctx,
		chromedp.Navigate(absLink),
		chromedp.Text("h1", &title, chromedp.ByQuery),
	)
	if err != nil {
		return "", err
	}
	return title, nil
}

func absPath(href *cdp.Node) string {
	return fmt.Sprintf(justwatchDomain + href.AttributeValue("href"))
}

func todayNetflixTag() string {
	timeNow := time.Now()
	return fmt.Sprintf("timeline__timeframe--%s--nfx", timeNow.Format(dateFormt))
}
