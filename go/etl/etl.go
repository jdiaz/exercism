// Package etl contains functions for facilitating data conversions between legacy and modern formats.
package etl

import "strings"

// Transform converts legacy data format into modern data format.
func Transform(legacy map[int][]string) map[string]int {
	modernDataMap := make(map[string]int)
	for key, letters := range legacy {
		for _, letter := range letters {
			modernDataMap[strings.ToLower(string(letter))] = key
		}
	}
	return modernDataMap
}
