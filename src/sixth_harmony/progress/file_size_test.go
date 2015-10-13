package progress

import (
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestFileSize(c *C) {
	filepath := "../../../test_short/dan.txt"
	size, _ := FileSize(filepath)

	c.Assert(size, Equals, int64(10))
}

func (s *MySuite) TestDirSize(c *C) {
	filepath := "../../../test"
	size, _ := DirSize(filepath)

	c.Assert(size, Equals, int64(4262))
}
