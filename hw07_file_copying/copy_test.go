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
		s.NoError(err)

		fromFile, err := os.Open(s.from)
		s.NoError(err)
		defer fromFile.Close()

		toFile, err := os.Open(s.to)
		s.NoError(err)
		defer toFile.Close()

		fromStat, err := fromFile.Stat()
		s.NoError(err)

		toStat, err := toFile.Stat()
		s.NoError(err)

		if s.limit == 0 {
			s.Equal(fromStat.Size()-s.offset, toStat.Size())
		} else {
			s.Equal(s.limit, toStat.Size()-s.offset)
		}
	})
}

func TestTSuite(t *testing.T) {
	suites := []*tSuite{
		{
			name: "copy ok",
			from: "testdata/input.txt",
			to:   "testdata/input_copy.txt",
		},
		{
			name:   "copy with offset",
			from:   "testdata/input.txt",
			to:     "testdata/input_copy_with_offset.txt",
			offset: 5,
		},
	}

	for i := 0; i < len(suites); i++ {
		suite.Run(t, suites[i])
	}
}
