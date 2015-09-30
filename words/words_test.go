package words

import (
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestCount(c *C) {
	wordCount := *Count("/Users/ofer987/go/src/dan/sixth_harmony/test_short")

	c.Assert(wordCount, Not(IsNil))

	c.Assert(wordCount["hello"], Equals, 2)
	c.Assert(wordCount["dan"], Equals, 1)
	c.Assert(wordCount["edith"], Equals, 1)
}
