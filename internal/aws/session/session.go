package session

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/candy12t/jarvis/internal/config"
)

var chachedSession *session.Session

func NewSession() *session.Session {
	if chachedSession != nil {
		return chachedSession
	}
	sess := session.Must(
		session.NewSession(
			&aws.Config{
				Credentials: credentials.NewStaticCredentials(config.AWSAccessKeyId(), config.AWSSecretAccessKey(), ""),
				Region:      aws.String(config.AWSRegion()),
			},
		),
	)
	chachedSession = sess
	return chachedSession
}
