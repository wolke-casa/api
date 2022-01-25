package s3

import (
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type Storage struct {
	bucket   string
	s3       *s3.S3
	uploader *s3manager.Uploader
}

func New(awsKey string, awsSecret string, awsRegion string, awsBucket string) *Storage {
	creds := credentials.NewStaticCredentials(awsKey, awsSecret, "")

	cfg := &aws.Config{
		Region:      aws.String(awsRegion),
		Credentials: creds,
		Endpoint:    aws.String("http://s3.wasabisys.com"),
	}

	session := session.Must(session.NewSession(cfg))

	return &Storage{
		s3:       s3.New(session),
		uploader: s3manager.NewUploader(session),
		bucket:   awsBucket,
	}
}

func (s *Storage) Put(content io.Reader, path string) error {
	_, err := s.uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(path),
		Body:   content,
	})

	return err
}

func (s *Storage) Get(path string) (io.ReadCloser, error) {
	resp, err := s.s3.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(path),
	})

	return resp.Body, err
}

func (s *Storage) Delete(path string) error {
	_, err := s.s3.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(path),
	})

	return err
}
