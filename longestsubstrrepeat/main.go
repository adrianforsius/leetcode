package main

import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 0 {
		return 0
	}

	wordList := make([]string, 0)
	longest := 1
	for _, char := range s {
		add := true
		for lineIndex, line := range wordList {
			if strings.ContainsRune(string(line), char) {
				if len(line) > longest {
					longest = len(line)
				}
				wordList[lineIndex] = string(char)
				add = false
			} else {
				line = line + string(char)
				wordList[lineIndex] = line
				if len(line) > longest {
					longest = len(line)
				}
			}
		}
		if add == true {
			wordList = append(wordList, string(char))
		}
		//fmt.Println("Add to list", wordList)
	}
	return longest
}

func main() {
	out := lengthOfLongestSubstring("au")
	fmt.Println("out", out)
}
