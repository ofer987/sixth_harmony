package sixth_harmony

import (
	"fmt"
	. "gopkg.in/check.v1"
)

func (s *MySuite) TestCountWordsInFile(c *C) {
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

	actualWordCount := *CountWordsInFile("/Users/ofer987/go/src/dan/sixth_harmony/test/good_morning.txt")

	c.Assert(compareMaps(actualWordCount, expectedWordCount), Equals, true)
}

func (s *MySuite) TestFindFilesWithAbsolutePath(c *C) {
	filepaths := []string{"/Users/ofer987/go/src/dan/sixth_harmony/test/.gitignore", "/Users/ofer987/go/src/dan/sixth_harmony/test/Gemfile", "/Users/ofer987/go/src/dan/sixth_harmony/test/Guardfile", "/Users/ofer987/go/src/dan/sixth_harmony/test/README.md", "/Users/ofer987/go/src/dan/sixth_harmony/test/good_morning.txt", "/Users/ofer987/go/src/dan/sixth_harmony/test/lib/sets.rb"}

	actualFilepaths := FindFiles("/Users/ofer987/go/src/dan/sixth_harmony/test")

	c.Assert(actualFilepaths, DeepEquals, filepaths)
}

func (s *MySuite) TestFindFilesWithRelativePath(c *C) {
	filepaths := []string{"../../test/.gitignore", "../../test/Gemfile", "../../test/Guardfile", "../../test/README.md", "../../test/good_morning.txt", "../../test/lib/sets.rb"}

	actualFilepaths := FindFiles("../../test")

	c.Assert(actualFilepaths, DeepEquals, filepaths)
}

func compareArrays(actual []string, expected []string) bool {
	if len(actual) != len(expected) {
		return false
	}

	for index, item := range actual {
		if item != expected[index] {
			fmt.Printf("%s\n%s\n", actual, expected)
			fmt.Printf("%s", item)
			fmt.Printf("%s", expected[index])
			return false
		}
	}

	return true
}
