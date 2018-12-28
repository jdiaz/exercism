// Package reverse contains functions to reverse strings.
package reverse

// String reverses a string.
func String(s string) string {
	arr := []rune(s)
	left := 0
	right := len(arr) - 1
	for left < right {
		tmp := arr[left]
		arr[left] = arr[right]
		arr[right] = tmp

		left++
		right--
	}
	return string(arr)
}
