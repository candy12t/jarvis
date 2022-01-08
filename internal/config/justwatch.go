package config

import (
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Justwatch struct {
	Country  string   `yaml:"country"`
	Services []string `yaml:"services"`
}

var justwachCfg *Justwatch

func JustwatchCfg() *Justwatch {
	return justwachCfg
}

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
