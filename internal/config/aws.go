package config

import (
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type AWSConfig struct {
	AccessKeyId     string   `yaml:"access_key_id"`
	SecretAccessKey string   `yaml:"secret_access_key"`
	Region          string   `yaml:"region"`
	S3              S3Config `yaml:"s3"`
}

type S3Config struct {
	Bucket string `yaml:"bucket"`
	Key    string `yaml:"key"`
}

var awsCfg *AWSConfig

func NewAWS() error {
	data, err := ReadFile(filepath.Join("/", "aws.yaml"))
	if err != nil {
		return err
	}

	cfg, err := ParseAWSConfig(data)
	if err != nil {
		return err
	}
	awsCfg = cfg
	return nil
}

func ParseAWSConfig(data []byte) (*AWSConfig, error) {
	aws := new(AWSConfig)
	if err := yaml.Unmarshal(data, aws); err != nil {
		return nil, err
	}
	return aws, nil
}

func AWSAccessKeyId() string {
	return awsCfg.AccessKeyId
}

func AWSSecretAccessKey() string {
	return awsCfg.SecretAccessKey
}

func AWSRegion() string {
	return awsCfg.Region
}

func AWSS3Bucket() string {
	return awsCfg.S3.Bucket
}

func AWSS3Key() string {
	return awsCfg.S3.Key
}
