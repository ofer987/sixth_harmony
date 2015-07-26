package calculator

import (
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestCompute(c *C) {
	wordCount := *Compute("test")

	c.Assert(wordCount, Not(IsNil))

	c.Assert(wordCount["hello"], Equals, 2)
	c.Assert(wordCount["dan"], Equals, 1)
	c.Assert(wordCount["edith"], Equals, 1)
}
