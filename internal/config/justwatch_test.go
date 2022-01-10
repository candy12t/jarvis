package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewJustwatch(t *testing.T) {
	err := NewJustwatch()
	assert.NoError(t, err)
	assert.Equal(t, "ja_JP", JustwatchCountry())
	assert.Equal(t, []string{"nfx", "amp", "dnp"}, JustwatchProviders())
}

func TestParseJustwatchConfig(t *testing.T) {
	data := `
country: ja_JP
providers:
  - nfx
  - amp
  - dnp
`

	cfg, err := ParseJustwatchConfig([]byte(data))
	assert.NoError(t, err)
	assert.Equal(t, "ja_JP", cfg.Country)
	assert.Equal(t, []string{"nfx", "amp", "dnp"}, cfg.Providers)
}
