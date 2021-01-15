package main

import (
	"fmt"
)

//func longestPalindrome(s string) string {
//        runes := []rune(s)
//        longest := string(runes[0])
//        for rIndex, r := range runes {
//                for ruIndex, ru := range runes {
//                        // Brute force this can be optmized to not
//                        // iterate over previously checked runes
//                        if rIndex <= ruIndex {
//                                continue
//                        }
//                        if r == ru {
//                                if ruIndex-rIndex > len(longest) {
//                                        candidate := runes[rIndex:ruIndex]
//                                        fmt.Println("candidate", candidate)
//                                        for i, j := 0, len(candidate); i < len(candidate)/2; i, j = i+i, j-j {
//                                                if candidate[i] != candidate[j] {
//                                                        break
//                                                }
//                                                longest = string(candidate)
//                                        }
//                                }
//                        }
//                }
//        }
//        return string(longest)
//}

func longestPalindrome(s string) string {
	if len(s) <= 0 {
		return ""
	}
	runes := []rune(s)
	longest := string(runes[0])
	for i := 0; i < len(runes); i++ {
		for j := len(runes) - 1; j > i/2; j-- {
			if runes[i] == runes[j] && i < j {
				candidate := runes[i : j+1]
				if len(candidate) > len(longest) {
					p := palindrome(candidate)
					if len(p) > len(longest) {
						longest = string(p)
					}
				}
			}
		}
	}
	return string(longest)
}

func palindrome(candidate []rune) []rune {
	for k, l := 0, len(candidate)-1; k <= len(candidate)/2; k, l = k+1, l-1 {
		if candidate[k] != candidate[l] {
			return []rune{}
		}
	}
	return candidate
}

func main() {
	out := longestPalindrome("abcda")
	fmt.Println("out", out)
}
