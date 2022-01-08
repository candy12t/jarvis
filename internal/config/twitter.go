package config

import (
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Twitter struct {
	ConsumerKey       string `yaml:"consumer_key"`
	ConsumerSecret    string `yaml:"consumer_secret"`
	AccessToken       string `yaml:"access_token"`
	AccessTokenSecret string `yaml:"access_token_secret"`
}

var twitterCfg *Twitter

func TwitterCfg() *Twitter {
	return twitterCfg
}

func NewTwitter() error {
	data, err := ReadFile(filepath.Join("/", "twitter.yaml"))
	if err != nil {
		return err
	}
	cfg, err := ParseTwitterConfig(data)
	if err != nil {
		return err
	}
	twitterCfg = cfg
	return nil
}

func ParseTwitterConfig(data []byte) (*Twitter, error) {
	t := new(Twitter)
	if err := yaml.Unmarshal(data, t); err != nil {
		return nil, err
	}
	return t, nil
}
