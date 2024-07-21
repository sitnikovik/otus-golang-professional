package main

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type tSuite struct {
	suite.Suite

	name   string
	from   string
	to     string
	offset int64
	limit  int64
	err    error
}

func (s *tSuite) SetupSuite() {
	removeFIleOnExists(s.to)
}

func (s *tSuite) TearDownSuite() {
	removeFIleOnExists(s.to)
}

func removeFIleOnExists(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return
	}

	if err := os.Remove(path); err != nil {
		log.Fatal(err)
	}
}

func (s *tSuite) TestCopy() {
	s.T().Run(s.name, func(t *testing.T) {
		err := Copy(s.from, s.to, s.offset, s.limit)
		if s.err != nil {
			s.Require().NoFileExists(s.to)
			s.Require().ErrorIs(err, s.err)
			return
		}

		s.Require().NoError(err)

		fromFile, err := os.Open(s.from)
		s.Require().NoError(err)
		defer fromFile.Close()

		toFile, err := os.Open(s.to)
		s.Require().NoError(err)
		defer toFile.Close()

		fromStat, err := fromFile.Stat()
		s.Require().NoError(err)

		toStat, err := toFile.Stat()
		s.Require().NoError(err)

		offset := s.offset
		if offset < 0 {
			offset = 0
		}
		if s.limit == 0 {
			s.Equal(fromStat.Size()-offset, toStat.Size())
		} else {
			s.IsNonIncreasing(s.limit, toStat.Size()-offset)
		}
	})
}

func TestTSuite(t *testing.T) {
	fromFile := "testdata/input.txt"
	suites := []*tSuite{
		{
			name: "copy ok",
			from: fromFile,
			to:   "testdata/input_copy.txt",
		},
		{
			name:   "err on offset less 0",
			from:   fromFile,
			to:     "testdata/input_copy_less0.txt",
			offset: -1,
			err:    ErrOffsetExceedsFileSize,
		},
		{
			name: "err unsported file on devurandom",
			from: "/dev/urandom",
			to:   "/tmp/213213123",
			err:  ErrUnsupportedFile,
		},
		{
			name: "err unsported file on dir",
			from: "testdata/dir",
			to:   "testdata/dir_copy",
			err:  ErrUnsupportedFile,
		},
		{
			name: "err on empty file",
			from: "testdata/empty.txt",
			to:   "testdata/empty_copy.txt",
			err:  ErrUnsupportedFile,
		},
		{
			name:  "offset 0 limit 100",
			from:  fromFile,
			to:    "testdata/out_offset0_limit100_copy.txt",
			limit: 100,
		},
		{
			name:  "offset 0 limit 1000",
			from:  fromFile,
			to:    "testdata/out_offset0_limit1000_copy.txt",
			limit: 1000,
		},
		{
			name:  "offset 0 limit 10000",
			from:  fromFile,
			to:    "testdata/out_offset0_limit10000_copy.txt",
			limit: 10000,
		},
		{
			name:   "offset 100 limit 1000",
			from:   fromFile,
			to:     "testdata/out_offset100_limit1000_copy.txt",
			offset: 100,
			limit:  1000,
		},
		{
			name:   "offset 6000 limit 10000",
			from:   fromFile,
			to:     "testdata/out_offset6000_limit10000_copy.txt",
			offset: 6000,
			limit:  10000,
		},
		{
			name:   "offset 6000 limit 10",
			from:   fromFile,
			to:     "testdata/out_offset6000_limit10_copy.txt",
			offset: 6000,
			limit:  10,
		},
	}

	for i := 0; i < len(suites); i++ {
		suite.Run(t, suites[i])
	}
}
