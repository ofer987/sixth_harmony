package main

import (
	"fmt"
	"os"
	"sixth_harmony"
)

func main() {
	currentDirectory, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error reading current directory")
		os.Exit(1)
	}
	fmt.Printf("Current Directory: %s\n", currentDirectory)

	wordCount := *sixth_harmony.CountWordsInDirectory(currentDirectory)
	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}
}

// func sortWords(map[string]int *words) map[string]int {
// 	sortedWords := make([string]int, len(words))
//
// 	for wordI, countI := range words {
// 		var tempCount int
// 		var tempWord string
//
// 		for wordJ, countJ := range words {
// 			if wordI == wordJ {
// 				continue
// 			}
//
// 			if countJ > countI {
// 				tempCount = countJ
// 				tempWord = wordJ
// 			}
// 		}
//
// 		// switch
// 		sortedWords
// 	}
// }
