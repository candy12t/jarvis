package config

import "os"

type Twitter struct {
	ComsumerKey       string
	ConmsumerSecret   string
	AccessToken       string
	AccessTokenSecret string
}

func NewTwitter() *Twitter {
	return &Twitter{
		ComsumerKey:       os.Getenv("ComsumerKey"),
		ConmsumerSecret:   os.Getenv("ComsumerSecret"),
		AccessToken:       os.Getenv("AccessToken"),
		AccessTokenSecret: os.Getenv("AccessTokenSecret"),
	}
}
