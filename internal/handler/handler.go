package handler

import (
	"time"

	"github.com/candy12t/jarvis/internal/config"
	"github.com/candy12t/jarvis/internal/fetcher"
	"github.com/candy12t/jarvis/internal/model/fetcher/justwatch"
	"github.com/candy12t/jarvis/internal/model/notifier/twitter"
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

	twitterAuth := twitter.NewAuthenticator(config.TwitterConsumerKey(), config.TwitterConsumerSecret(), config.TwitterAccessToken(), config.TwitterAccessTokenSecret())
	notifierClient := twitter.NewClient(twitterAuth)
	// notifierClient := dumynotifier.NewClient()
	_notifier := notifier.NewClient(notifierClient, timeNow)

	return _notifier.Notify(contents)
}
