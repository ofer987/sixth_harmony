package calculator

import (
	"dan/sixth_harmony/file"
)

func Compute(directory string) *map[string]int {
	filepaths := file.FindFiles("../test_short")

	masterWordCount := make(map[string]int)
	var wordCount map[string]int
	for _, filepath := range filepaths {
		wordCount = *file.EvaluateFile("../test_short" + "/" + filepath)

		for word, count := range wordCount {
			if masterCount, ok := masterWordCount[word]; ok {
				masterWordCount[word] = masterCount + count
			} else {
				masterWordCount[word] = count
			}
		}
	}

	return &masterWordCount
}
