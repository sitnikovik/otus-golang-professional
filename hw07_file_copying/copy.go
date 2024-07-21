package main

import (
	"errors"
	"io"
	"os"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	if offset < 0 {
		return ErrOffsetExceedsFileSize
	}

	fromFile, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer fromFile.Close()

	toFile, err := os.Create(toPath)
	if err != nil {
		return err
	}

	stat, err := fromFile.Stat()
	if err != nil {
		return err
	}
	if stat.IsDir() {
		return ErrUnsupportedFile
	}

	if offset > stat.Size() {
		return ErrOffsetExceedsFileSize
	}

	if limit == 0 {
		limit = stat.Size() - offset
	}

	if _, err := fromFile.Seek(offset, 0); err != nil {
		return err
	}

	if _, err := io.CopyN(toFile, fromFile, limit); err != nil {
		if err == io.EOF {
			return nil
		}
		return err
	}

	if err := toFile.Close(); err != nil {
		return err
	}

	return nil
}
