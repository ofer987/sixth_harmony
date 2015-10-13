package progress

import (
	"io/ioutil"
	"os"
)

func FileSize(filepath string) (int64, error) {
	file, err := os.Open(filepath)
	defer file.Close()

	if err != nil {
		return int64(0), err
	}

	fi, err := file.Stat()
	if err != nil {
		return int64(0), err
	}

	return fi.Size(), nil
}

func DirSize(filepath string) (int64, error) {
	// Ignore hidden files
	if filepath[0] == '.' {
		return 0, nil
	}

	FileSize(filepath)
	size := int64(0)

	files, err := ioutil.ReadDir(filepath)
	if err != nil {
		return int64(0), err
	}

	for _, file := range files {
		if file.IsDir() {
			dirSize, err := DirSize(filepath + "/" + file.Name())
			if err != nil {
				return int64(0), err
			}
			size = size + dirSize
		} else if file.Name()[0] == '.' {
			// Ignore hidden files
		}

		size = size + file.Size()
	}

	return size, nil
}
