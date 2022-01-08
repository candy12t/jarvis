package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseJustwatchConfig(t *testing.T) {
	data := `
country: jp_JP
services:
  - netflix
  - amazon-prime-video
  - disney-plus
`

	cfg, err := ParseJustwatchConfig([]byte(data))
	assert.NoError(t, err)
	assert.Equal(t, "jp_JP", cfg.Country)
	assert.Equal(t, []string{"netflix", "amazon-prime-video", "disney-plus"}, cfg.Services)
}
