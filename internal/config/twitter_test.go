package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseTwitterConfig(t *testing.T) {
	data := `
consumer_key: INPUT_YOUR_CONSUMER_KEY
consumer_secret: INPUT_YOUR_CONSUMER_SECRET
access_token: INPUT_YOUR_ACCESS_TOKEN
access_token_secret: INPUT_YOUR_ACCESS_TOKEN_SECRET
`

	cfg, err := ParseTwitterConfig([]byte(data))
	assert.NoError(t, err)
	assert.Equal(t, "INPUT_YOUR_CONSUMER_KEY", cfg.ConsumerKey)
	assert.Equal(t, "INPUT_YOUR_CONSUMER_SECRET", cfg.ConsumerSecret)
	assert.Equal(t, "INPUT_YOUR_ACCESS_TOKEN", cfg.AccessToken)
	assert.Equal(t, "INPUT_YOUR_ACCESS_TOKEN_SECRET", cfg.AccessTokenSecret)
}
