package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency docs
func ConcurrentFrequency(texts []string) FreqMap {
	res := FreqMap{}
	mapChan := make(chan FreqMap, len(texts))

	for _, text := range texts {
		go func(s string) {
			mapChan <- Frequency(s)
		}(text)
	}

	// Loop over the channel and combining the results
	for range texts {
		for letter, freq := range <-mapChan {
			res[letter] += freq
		}
	}
	close(mapChan)
	return res
}
