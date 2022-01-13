package handler

import (
	"time"

	"github.com/candy12t/jarvis/internal/config"
	"github.com/candy12t/jarvis/internal/fetcher"
	"github.com/candy12t/jarvis/internal/fetcher/justwatch"
	"github.com/candy12t/jarvis/internal/notifier"
)

const dateFormat = "2006-01-02"

func Apply() error {
	timeNow := time.Now().Format(dateFormat)

	fetcherClient := justwatch.NewClient(config.JustwatchCountry())
	_fetcher := fetcher.NewClient(fetcherClient, timeNow)

	contents, err := _fetcher.FetchContents()
	if err != nil {
		return err
	}

	_notifier := notifier.NewClient(NotifyClient(), timeNow)

	return _notifier.Notify(contents)
}
