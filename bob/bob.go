// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

// Hey responds to a greeting like a lackadaisical teenager
func Hey(greeting string) string {
	// Write some code here to pass the test suite.
	greeting = strings.TrimSpace(greeting)

	switch {
	case greeting == "":
		return "Fine. Be that way!"

	case greeting[len(greeting)-1] == '?' && any(greeting, unicode.IsLetter) && !any(greeting, unicode.IsLower):
		return "Calm down, I know what I'm doing!"

	case greeting[len(greeting)-1] == '?':
		return "Sure."

	case any(greeting, unicode.IsLetter) && !any(greeting, unicode.IsLower):
		return "Whoa, chill out!"

	default:
		return "Whatever."
	}
}

func any(str string, testFunc func(rune) bool) bool {
	for _, c := range str {
		if testFunc(c) {
			return true
		}
	}
	return false
}
