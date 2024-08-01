package main

import (
	"errors"
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

	// Open the file to copy
	fromFile, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer fromFile.Close()

	// Validate
	stat, err := fromFile.Stat()
	if err != nil || stat.IsDir() || stat.Size() == 0 {
		return ErrUnsupportedFile
	}
	if offset > stat.Size() {
		return ErrOffsetExceedsFileSize
	}

	// Start copying
	toFile, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer toFile.Close()
	if limit == 0 || limit > stat.Size() {
		limit = stat.Size()
	}
	return runCopy(fromFile, toFile, offset, limit)
}

func runCopy(fromFile *os.File, toFile *os.File, offset, limit int64) error {
	if _, err := fromFile.Seek(offset, 0); err != nil {
		return err
	}

	// start new bar
	bar := pb.Full.Start64(limit)
	defer bar.Finish()

	// create proxy reader
	barReader := bar.NewProxyReader(fromFile)

	if _, err := io.CopyN(toFile, barReader, limit); err != nil {
		if !errors.Is(err, io.EOF) {
			return err
		}
	}

	return nil
}
