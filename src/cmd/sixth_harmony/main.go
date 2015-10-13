package main

import (
	"fmt"
	"os"
	"sixth_harmony"
	"sixth_harmony/progress"
)

func main() {
	specifiedDir := directory()
	fmt.Printf("Current Directory: %s\n", specifiedDir)

	size, err := progress.DirSize(specifiedDir)
	if err != nil {
		os.Exit(1)
	}

	fmt.Printf("Size: %v\n", size)

	wordCount := *sixth_harmony.CountWordsInDirectory(specifiedDir)
	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}
}

func directory() string {
	specified := os.Getenv("dir")
	if specified == "" {
		current, err := os.Getwd()
		if err != nil {
			return "./"
		}

		return current
	}

	return specified
}
