package grains

import "fmt"

// Square docs
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, fmt.Errorf("error")
	}
	return 1 << uint(n-1), nil
}

// Total docs
func Total() (t uint64) {
	return 1<<64 - 1
}
