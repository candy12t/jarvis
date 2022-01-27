package model

type WrappedContents struct {
	Contents Contents `json:"contents,omitempty"`
}

type Contents []Content

type Content struct {
	Title      string `json:"title"`
	ContentURL string `json:"content_url"`
}
