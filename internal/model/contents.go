package model

type FetchResult struct {
	Contents Contents `json:"contents"`
}

type Contents []Content

type Content struct {
	Title      string `json:"title"`
	ContentURL string `json:"content_url"`
}
