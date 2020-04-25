package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram docs
func IsIsogram(input string) bool {
	seen := map[rune]bool{}

	for _, char := range strings.ToUpper(input) {
		if !unicode.IsLetter(char) {
			continue
		}
		_, ok := seen[char]
		if ok {
			return false
		}
		seen[char] = true
	}
	return true
}

// IsIsogram Hung Solution
/*func IsIsogram(input string) bool {
	var check = make(map[rune]bool)
	for i := 'A'; i <= 'Z'; i++ {
		check[i] = false
	}

	var isIsogram bool = true
	for _, c := range strings.ToUpper(input) {
		if !unicode.IsLetter(c) {
			continue
		} else if check[c] == false {
			check[c] = true
		} else {
			isIsogram = false
		}
	}
	return isIsogram
}*/
