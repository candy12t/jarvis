package fetcher

import "github.com/candy12t/jarvis/internal/model"

type Fetcher interface {
	FetchContents(string) (model.Contents, error)
}

type Fetch struct {
	Fetcher
	date string
}

func NewClient(fetcher Fetcher, date string) *Fetch {
	return &Fetch{
		Fetcher: fetcher,
		date:    date,
	}
}

func (f *Fetch) FetchContents() (model.Contents, error) {
	return f.Fetcher.FetchContents(f.date)
}
