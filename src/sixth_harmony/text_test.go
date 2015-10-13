package sixth_harmony

import (
	"fmt"
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

// func (s *MySuite) Test_readFile(c *C) {
// 	contents := readFile("/Users/ofer987/go/src/dan/sixth_harmony/test/good_morning.txt")
//
// 	c.Assert(contents, Equals, "Hello Dan, lets read the book Good Morning Canada together please.\nDan read Good Night Canada\nhello\n")
// }

func (s *MySuite) TestcountWords(c *C) {
	expectedWordCount := make(map[string]int)

	expectedWordCount["hello"] = 2
	expectedWordCount["dan"] = 2
	expectedWordCount["lets"] = 1
	expectedWordCount["read"] = 2
	expectedWordCount["the"] = 1
	expectedWordCount["book"] = 1
	expectedWordCount["good"] = 2
	expectedWordCount["morning"] = 1
	expectedWordCount["canada"] = 2
	expectedWordCount["together"] = 1
	expectedWordCount["please"] = 1
	expectedWordCount["night"] = 1

	actualContent := "hello dan lets read the book good morning canada together please dan read good night canada hello"
	actualWordCount := *countWords(actualContent)

	c.Assert(compareMaps(actualWordCount, expectedWordCount), Equals, true)
}

func (s *MySuite) TestAll(c *C) {
	content := "The code was:\nh128\nIt was used 12 times."

	expectedWordCount := map[string]int{
		"the":   1,
		"code":  1,
		"was":   2,
		"it":    1,
		"used":  1,
		"times": 1,
	}

	actualWordCount := *countWords(removeNonWordChars(content))

	c.Assert(compareMaps(actualWordCount, expectedWordCount), Equals, true)
}

func (s *MySuite) TestonlyOneWhiteSpace(c *C) {
	rawString := "  remove \nOk?"
	expectedString := "remove ok"

	c.Assert(removeNonWordChars(rawString), Equals, expectedString)
}

func (s *MySuite) TestremoveNonWordChars(c *C) {
	rawString := "Dan, please remove all non-word chars!\nOk?"
	expectedString := "dan please remove all non word chars ok"

	c.Assert(removeNonWordChars(rawString), Equals, expectedString)

	rawString = "D2n"
	expectedString = "d2n"

	c.Assert(removeNonWordChars(rawString), Equals, expectedString)
}

func compareMaps(actual map[string]int, expected map[string]int) bool {
	if len(actual) != len(expected) {
		fmt.Printf("%s\n%s", actual, expected)
		return false
	}

	for key, value := range actual {
		if expectedValue, ok := expected[key]; ok {
			if expectedValue != value {
				fmt.Printf("Actual: %s\nExpected: %s", actual, expected)
				return false
			}
		}
	}

	return true
}
