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

	// Make files to work with
	fromFile, toFile, err := makeFiles(fromPath, toPath)
	if err != nil {
		return err
	}
	defer func() {
		fromFile.Close()
		toFile.Close()
	}()

	// Validate file to copy
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

	// Start copying
	if _, err := fromFile.Seek(offset, 0); err != nil {
		return err
	}
	if limit == 0 {
		limit = stat.Size()
	}
	if _, err := io.CopyN(toFile, fromFile, limit); err != nil {
		if err == io.EOF {
			return nil
		}
		return err
	}

	return nil
}

func makeFiles(fromPath, toPath string) (*os.File, *os.File, error) {
	fromFile, err := os.Open(fromPath)
	if err != nil {
		return nil, nil, err
	}

	toFile, err := os.Create(toPath)
	if err != nil {
		return nil, nil, err
	}

	return fromFile, toFile, nil
}
