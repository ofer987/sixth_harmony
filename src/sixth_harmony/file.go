package sixth_harmony

import (
	"fmt"
	"io/ioutil"
	"sixth_harmony/progress"
)

func CountWordsInDirectory(directory string) *map[string]int {
	filepaths := FindFiles(directory)

	masterWordCount := make(map[string]int)
	var wordCount map[string]int
	for _, filepath := range filepaths {
		wordCount = *CountWordsInFile(filepath)

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

func CountWordsInFile(filepath string) *map[string]int {
	content := readFile(filepath)
	content = removeNonWordChars(content)

	return countWords(content)
}

func FindFiles(directory string) []string {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		return []string{}
	}

	filenames := []string{}
	for _, file := range files {
		if file.IsDir() {
			nestedDirectory := directory + "/" + file.Name()
			nestedFiles := FindFiles(nestedDirectory)
			for _, nestedFile := range nestedFiles {
				filenames = append(filenames, nestedFile)
			}
		} else {
			filenames = append(filenames, directory+"/"+file.Name())
		}
	}
	return filenames
}

func readFile(filepath string) string {
	// Ignore hidden files
	if filepath[0] == '.' {
		return ""
	}

	content, err := ioutil.ReadFile(filepath)

	size, err := progress.FileSize(filepath)
	if err != nil {
		return ""
	}
	fmt.Printf("%v:\t%v\n", filepath, size)

	if err != nil {
		return ""
	}

	return string(content)
}
