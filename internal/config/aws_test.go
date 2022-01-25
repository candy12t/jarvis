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
	assert.Equal(t, "INPUT_YOUR_S3_KEY", AWSS3Key())
}

func TestParseAWSConfig(t *testing.T) {
	data := `
access_key_id: INPUT_YOUR_ACCESS_KEY_ID
secret_access_key: INPUT_YOUR_SECRET_ACCESS_KEY
region: INPUT_YOUR_REGION
s3:
  bucket: INPUT_YOUR_S3_BUCKET
  key: INPUT_YOUR_S3_KEY
`

	cfg, err := ParseAWSConfig([]byte(data))
	assert.NoError(t, err)
	assert.Equal(t, "INPUT_YOUR_ACCESS_KEY_ID", cfg.AccessKeyId)
	assert.Equal(t, "INPUT_YOUR_SECRET_ACCESS_KEY", cfg.SecretAccessKey)
	assert.Equal(t, "INPUT_YOUR_REGION", cfg.Region)
	assert.Equal(t, "INPUT_YOUR_S3_BUCKET", cfg.S3.Bucket)
	assert.Equal(t, "INPUT_YOUR_S3_KEY", cfg.S3.Key)
}
