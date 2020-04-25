package raindrops

import "strconv"

// Convert doc
func Convert(input int) string {
	var sound string

	if input%3 == 0 {
		sound += "Pling"
	}
	if input%5 == 0 {
		sound += "Plang"
	}
	if input%7 == 0 {
		sound += "Plong"
	}
	if sound == "" {
		sound = strconv.Itoa(input)
	}
	return sound
}
