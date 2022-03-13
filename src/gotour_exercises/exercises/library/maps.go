package library

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	count := make(map[string]int)

	for _, word := range strings.Fields(s) {
		_, ok := count[word]
		if ok {
			count[word] += 1
		} else {
			count[word] = 1
		}
	}

	return count
}

func RunWordCount() {
	wc.Test(WordCount)
}
