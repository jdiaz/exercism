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

// ConcurrentFrequency counts the frequency of each rune in a given text concurrently
// and returns thos data as a FreqMap
func ConcurrentFrequency(texts []string) FreqMap {
	n := len(texts)
	c := make(chan FreqMap, n)
	for _, text := range texts {
		go func(s string) {
			c <- Frequency(s)
		}(text)
	}

	out := FreqMap{}
	for i := 0; i < n; i++ {
		tmpFreqmap := <-c
		for k, v := range tmpFreqmap {
			out[k] = out[k] + v
		}
	}

	return out
}
