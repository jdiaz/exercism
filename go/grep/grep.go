// Package grep contains functions for finding files in the filesystem.
package grep

import (
	"bufio"
	"log"
	"os"
	"strings"
	"strconv"
)

const (
	N = "-n"
	L = "-l"
	I = "-i"
	V = "-v"
	X = "-x"
)

// Search find files who's contents matches the supplied pattern.
func Search(pattern string, flags, files []string) []string {
	flagProvidedMap := make(map[string]bool, len(flags))
	for _, flag := range flags {
		flagProvidedMap[flag] = true
	}
	return lineMatchSearch(pattern, files, flagProvidedMap)
}

func lineMatchSearch(pattern string, files []string, flagMap map[string]bool) []string {
	matches := make([]string, 0)
	for _, filename := range files {
		f, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
			return matches
		}
		scanner := bufio.NewScanner(f)
		lineNumber := 1
		for scanner.Scan() {
			line := scanner.Text()
			match := false
			match = strings.Contains(line, pattern)
			if flagMap[I] && !flagMap[X] {
				// Case insenstive
				match = strings.Contains(strings.ToLower(line), strings.ToLower(pattern))
			}
			if flagMap[X] && !flagMap[I] {
				match = line == pattern
			}
			if flagMap[X] && flagMap[I] {
				match = strings.ToLower(line) == strings.ToLower(pattern)
			}
			if match && !flagMap[V] {
				newLine := line
				if flagMap[N] {
					newLine = strconv.Itoa(lineNumber) + ":" + line
				}
				if flagMap[L] {
					newLine = filename
					//matches = append(matches, newLine)
					//break
				}
				matches = append(matches, newLine)
			}
			if !match && flagMap[V] {
				matches = append(matches, line)
			}
			lineNumber++
		}
		f.Close()
	}

	return matches
}
