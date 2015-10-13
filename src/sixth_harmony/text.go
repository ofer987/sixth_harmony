package sixth_harmony

import (
	"regexp"
	"strings"
)

func removeNonWordChars(rawString string) string {
	lowerCaseString := strings.ToLower(rawString)

	// Remove non-word characters
	notWordRegex := regexp.MustCompile("\\W")
	replacementBytes := notWordRegex.ReplaceAll([]byte(lowerCaseString), []byte{' '})
	replacementString := string(replacementBytes)

	// Remove leading and trailing whitespace
	replacementString = strings.Trim(replacementString, " ")

	// Remove double whitespace
	whitespaceRegex := regexp.MustCompile("\\s+")
	replacementBytes = whitespaceRegex.ReplaceAll([]byte(replacementString), []byte{' '})
	replacementString = string(replacementBytes)

	return replacementString
}

func countWords(content string) *map[string]int {
	words := strings.Split(content, " ")

	wordCount := make(map[string]int)
	for _, word := range words {
		if isWord(word) {
			if count, ok := wordCount[word]; ok {
				wordCount[word] = count + 1
			} else {
				wordCount[word] = 1
			}
		}
	}

	return &wordCount
}

func isWord(str string) bool {
	regex := regexp.MustCompile("^[a-zA-Z_-]+$")

	return regex.MatchString(str)
}
