package luhn

import (
	"strings"
)

// Valid docs
func Valid(input string) bool {
	// remove spaces
	s := strings.ReplaceAll(input, " ", "")

	// check if length is bigger than 1 or not
	if len(s) <= 1 {
		return false
	}

	// Modified SIN and sum up
	var sum int
	for i, c := range s {
		// Convert rune to int and check if is number or not
		num := int(c - '0')
		if num < 0 || num > 9 {
			return false
		}
		// Only modified every second digit from the right
		if (len(s)-i)%2 != 0 {
			sum += num
		} else {
			if num*2 > 9 {
				sum += num*2 - 9
			} else {
				sum += num * 2
			}
		}
	}
	return sum%10 == 0
}

// remove spaces
//s := strings.Join(strings.Fields(str), "")

// Convert rune to int and check if is number or not
// num, err := strconv.Atoi(string(c))
// if err != nil {
// 	return false
// }
