package main

import (
	"log"

	"github.com/candy12t/jarvis/internal/config"
	"github.com/candy12t/jarvis/internal/controller"
	"github.com/candy12t/jarvis/internal/twitter"
)

func main() {
	cfg := config.NewTwitter()
	client, err := twitter.NewClient(cfg.ComsumerKey, cfg.ConmsumerSecret, cfg.AccessToken, cfg.AccessTokenSecret)
	if err != nil {
		log.Fatal(err)
	}
	if err := controller.TweetNewMedia(*client); err != nil {
		log.Fatal(err)
	}
}
