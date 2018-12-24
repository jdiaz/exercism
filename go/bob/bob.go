// Package bob contains functions for responding to teenager Bob.
package bob

import (
	"strings"
	"unicode"
)

var categoryResponseMap = map[string]string{
	"question":      "Sure.",
	"yell":          "Whoa, chill out!",
	"questionshout": "Calm down, I know what I'm doing!",
	"silence":       "Fine. Be that way!",
	"anything":      "Whatever.",
}

// Hey function provides answers to interactions with Bob.
func Hey(remark string) string {
	return categoryResponseMap[getRemarkCategory(remark)]
}

func getRemarkCategory(remark string) string {

	switch r := strings.TrimSpace(remark); {
	case isRemarkQuestion(r) && isRemarkAllInCaps(r):
		return "questionshout"
	case isRemarkSilence(r):
		return "silence"
	case isRemarkQuestion(r):
		return "question"
	case isRemarkAllInCaps(r):
		return "yell"
	}
	return "anything"
}

func isRemarkSilence(remark string) bool {
	return remark == ""
}

func isRemarkQuestion(remark string) bool {
	n := len(remark)
	if n == 0 {
		return false
	}
	return remark[n-1:] == "?"
}

func isRemarkAllInCaps(remark string) bool {
	letterCount := 0
	for _, char := range remark {
		if !unicode.IsLetter(char) {
			continue
		}
		if !unicode.IsUpper(char) {
			return false
		}
		letterCount++
	}
	return letterCount > 0
}
