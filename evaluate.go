package file

import (
	"io/ioutil"
)

func EvaluateFile(filepath string) *map[string]int {
	content := readFile(filepath)
	content = toValidString(content)

	return wordCount(content)
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
				filenames = append(filenames, file.Name()+"/"+nestedFile)
			}
		} else {
			filenames = append(filenames, file.Name())
		}
	}
	return filenames
}
