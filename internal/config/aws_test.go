package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAWS(t *testing.T) {
	err := NewAWS()
	assert.NoError(t, err)
	assert.Equal(t, "INPUT_YOUR_ACCESS_KEY_ID", AWSAccessKeyId())
	assert.Equal(t, "INPUT_YOUR_SECRET_ACCESS_KEY", AWSSecretAccessKey())
	assert.Equal(t, "INPUT_YOUR_REGION", AWSRegion())
	assert.Equal(t, "INPUT_YOUR_S3_BUCKET", AWSS3Bucket())
}

func TestParseAWSConfig(t *testing.T) {
	data := `
access_key_id: INPUT_YOUR_ACCESS_KEY_ID
secret_access_key: INPUT_YOUR_SECRET_ACCESS_KEY
region: INPUT_YOUR_REGION
s3:
  bucket: INPUT_YOUR_S3_BUCKET
`

	cfg, err := ParseAWSConfig([]byte(data))
	assert.NoError(t, err)
	assert.Equal(t, "INPUT_YOUR_ACCESS_KEY_ID", cfg.AccessKeyId)
	assert.Equal(t, "INPUT_YOUR_SECRET_ACCESS_KEY", cfg.SecretAccessKey)
	assert.Equal(t, "INPUT_YOUR_REGION", cfg.Region)
	assert.Equal(t, "INPUT_YOUR_S3_BUCKET", cfg.S3.Bucket)
}
