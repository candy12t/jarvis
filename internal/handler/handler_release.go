//go:build !debug

package handler

import (
	"github.com/candy12t/jarvis/internal/config"
	"github.com/candy12t/jarvis/internal/notifier"
	"github.com/candy12t/jarvis/internal/notifier/twitter"
)

func NotifyClient() notifier.Notifier {
	twitterAuth := twitter.NewAuthenticator(
		config.TwitterConsumerKey(),
		config.TwitterConsumerSecret(),
		config.TwitterAccessToken(),
		config.TwitterAccessTokenSecret(),
	)
	return twitter.NewClient(twitterAuth)
}
