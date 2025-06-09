package utils

import (
	"strconv"
)

func AddPrefixSuffix(arr, prefixes, suffixes []string) []string {
	res := make([]string, 0, len(arr))

	for idx, word := range arr {
		var prefix, suffix string
		if idx < len(prefixes) {
			prefix = prefixes[idx]
		}
		if idx < len(suffixes) {
			suffix = suffixes[idx]
		}

		res = append(res, prefix+word+suffix)
	}

	return res
}

func NumbersAsStrings(start, end int) []string {
	res := make([]string, 0, end-start+1)

	for i := start; i <= end; i++ {
		res = append(res, strconv.Itoa(i))
	}

	return res
}

func RepeatWord(word string, n int) []string {
	res := make([]string, 0, n)

	for range n {
		res = append(res, word)
	}

	return res
}
