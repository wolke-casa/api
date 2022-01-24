package local

import (
	"io"
	"os"
	"path/filepath"
)

type Storage struct {
	Base string
}

func New(base string) *Storage {
	return &Storage{Base: base}
}

func (s *Storage) Put(content io.Reader, path string) error {
	out, err := os.Create(filepath.Join(s.Base, path))
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, content)

	return err
}

func (s *Storage) Get(path string) (io.ReadCloser, error) {
	file, err := os.Open(filepath.Join(s.Base, path))
	if err != nil {
		return nil, err
	}

	return file, err
}

func (s *Storage) Delete(path string) error {
	return os.Remove(filepath.Join(s.Base, path))
}
