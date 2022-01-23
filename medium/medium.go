package medium

import (
	"errors"
	"io"

	"github.com/wolke-gallery/api/config"
	"github.com/wolke-gallery/api/medium/local"
	"github.com/wolke-gallery/api/medium/s3"
)

type Medium interface {
	Put(content io.Reader, path string) error
	Get(path string) (io.ReadCloser, error)
	Delete(path string) error
}

var Storage Medium

func Initialize() error {
	switch config.Config.Medium {
	case "S3":
		s3Storage := s3.New(config.Config.AwsKey, config.Config.AwsSecret, config.Config.AwsRegion, config.Config.AwsBucket)
		Storage = s3Storage
	case "local":
		localStorage := local.New(config.Config.Directory)
		Storage = localStorage
	default:
		return errors.New("invalid `medium`. must be `S3` or `local`")
	}

	return nil
}
