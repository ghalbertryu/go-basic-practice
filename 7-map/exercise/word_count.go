package main

import (
	"golang.org/x/tour/wc"
	"strings"
	"unicode"
)

func main() {
	wc.Test(WordCount)
}

func WordCount(s string) map[string]int {
	wordSlice := strings.FieldsFunc(s, func(c rune) bool {
		return unicode.IsSpace(c)
	})

	ret := make(map[string]int)
	for _, v := range wordSlice {
		_, exist := ret[v]
		if exist {
			ret[v] += 1
		} else {
			ret[v] = 1
		}
	}

	return ret
}
