package storage

import (
	"errors"
	"io"

	"github.com/wolke-gallery/api/config"
	"github.com/wolke-gallery/api/storage/local"
	"github.com/wolke-gallery/api/storage/s3"
)

type storage interface {
	Put(content io.Reader, path string) error
	Get(path string) (io.ReadCloser, error)
	Delete(path string) error
}

var Storage storage

func Initialize() error {
	switch config.Config.Storage {
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
