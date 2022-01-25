package s3

import (
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3Service struct {
	Bucket     string
	Key        string
	uploader   *s3manager.Uploader
	downloader *s3manager.Downloader
}

func NewS3Service(sess *session.Session, bucket, key string) *S3Service {
	return &S3Service{
		Bucket:     bucket,
		Key:        key,
		uploader:   s3manager.NewUploader(sess),
		downloader: s3manager.NewDownloader(sess),
	}
}

func (s *S3Service) Upload(body io.Reader) error {
	_, err := s.uploader.Upload(
		&s3manager.UploadInput{
			Bucket: aws.String(s.Bucket),
			Key:    aws.String(s.Key),
			Body:   body,
		},
	)
	return err
}

func (s *S3Service) Download() ([]byte, error) {
	buf := aws.NewWriteAtBuffer([]byte{})
	_, err := s.downloader.Download(
		buf,
		&s3.GetObjectInput{
			Bucket: aws.String(s.Bucket),
			Key:    aws.String(s.Key),
		},
	)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
