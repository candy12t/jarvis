package config

import (
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Justwatch struct {
	Country   string   `yaml:"country"`
	Providers []string `yaml:"providers"`
}

var justwachCfg *Justwatch

func NewJustwatch() error {
	data, err := ReadFile(filepath.Join("/", "justwatch.yaml"))
	if err != nil {
		return err
	}

	cfg, err := ParseJustwatchConfig(data)
	if err != nil {
		return err
	}
	justwachCfg = cfg
	return nil
}

func ParseJustwatchConfig(data []byte) (*Justwatch, error) {
	jw := new(Justwatch)
	if err := yaml.Unmarshal(data, jw); err != nil {
		return nil, err
	}
	return jw, nil
}

func JustwatchCountry() string {
	return justwachCfg.Country
}

func JustwatchProviders() []string {
	return justwachCfg.Providers
}
