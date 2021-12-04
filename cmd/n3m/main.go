package main

import (
	"log"

	"github.com/candy12t/n3m/internal/config"
	"github.com/candy12t/n3m/internal/controller"
	"github.com/candy12t/n3m/internal/twitter"
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
