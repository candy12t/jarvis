package session

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

var chachedSession *session.Session

func NewSession(access_key_id, secret_access_key, region string) *session.Session {
	sess := session.Must(
		session.NewSession(
			&aws.Config{
				Credentials: credentials.NewStaticCredentials(access_key_id, secret_access_key, ""),
				Region:      aws.String(region),
			},
		),
	)
	chachedSession = sess
	return chachedSession
}
