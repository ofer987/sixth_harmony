package sixth_harmony

import (
	"io/ioutil"
)

func CountWordsInFile(filepath string) *map[string]int {
	content := readFile(filepath)
	content = RemoveNonWordChars(content)

	return CountWords(content)
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
	content, err := ioutil.ReadFile(filepath)

	if err != nil {
		return ""
	}

	return string(content)
}
