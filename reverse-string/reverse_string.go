package reverse

// Reverse docs
func Reverse(input string) string {
	inputRune := []rune(input)
	l := len(inputRune)

	result := make([]rune, l)
	for _, c := range inputRune {
		l--
		result[l] = c
	}
	return string(result)
	// res := ""
	// for _, v := range s {
	// 	res = string(v) + res

	// }
	// return res
}
