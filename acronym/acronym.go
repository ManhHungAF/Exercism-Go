// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	abbreviateArray := []string{}
	takeLetter := true
	for _, runes := range s {
		if takeLetter == true && unicode.IsLetter(runes) {
			abbreviateArray = append(abbreviateArray, strings.ToUpper(string(runes)))
		}

		if unicode.IsSpace(runes) || (unicode.IsPunct(runes) && string(runes) != "'") {
			takeLetter = true
		} else {
			takeLetter = false
		}
	}

	return strings.Join(abbreviateArray, "")
}

// Solution of Hung
// AbbreviateHung should have a comment documenting it.
/*func AbbreviateHung(str string) string {
	var results []string
	// Split phrase to seperate words
	words := strings.Split(str, " ")
	for i, word := range words {
		t := strings.Split(word, "-")
		if len(t) > 1 {
			words = append(words[:i], append(t, words[i+1:]...)...)
		}
	}

	for _, word := range words {
		wordRune := []rune(word)
		var firstLetter string

		for _, c := range wordRune {
			if unicode.IsLetter(c) {
				firstLetter = strings.ToUpper(string(c))
				break
			}
		}
		results = append(results, firstLetter)
	}

	return strings.Join(results, "")
}*/
