package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
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
	if err != nil || stat.IsDir() || stat.Size() == 0 {
		return ErrUnsupportedFile
	}

	if offset > stat.Size() {
		return ErrOffsetExceedsFileSize
	}

	// Start copying
	if limit == 0 || limit > stat.Size() {
		limit = stat.Size()
	}
	return runCopy(fromFile, toFile, offset, limit)
}

func makeFiles(fromPath, toPath string) (*os.File, *os.File, error) {
	fromFile, err := os.Open(fromPath)
	if err != nil {
		return nil, nil, ErrUnsupportedFile
	}

	toFile, err := os.Create(toPath)
	if err != nil {
		return nil, nil, err
	}

	return fromFile, toFile, nil
}

func runCopy(fromFile *os.File, toFile *os.File, offset, limit int64) error {
	if _, err := fromFile.Seek(offset, 0); err != nil {
		return err
	}

	// start new bar
	fmt.Printf("o: %d, l: %d\n", offset, limit)
	bar := pb.Full.Start64(limit)
	defer bar.Finish()

	// create proxy reader
	barReader := bar.NewProxyReader(fromFile)

	if _, err := io.CopyN(toFile, barReader, limit); err != nil {
		if err == io.EOF {
			return nil
		}
		return err
	}

	return nil
}
