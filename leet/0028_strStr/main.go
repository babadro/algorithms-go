package main

import "fmt"

func main() {
	fmt.Println(strStr("mississippi", "issip"))
}

func strStr(haystack string, needle string) int {
	needleLen := len(needle)
	if needleLen == 0 {
		return 0
	}
	haystackLen := len(haystack)
	if haystackLen == 0 {
		return -1
	}

	i := 0
	j := 0
	for i < haystackLen {
		if haystack[i] == needle[j] {
			j++
			i++
		} else {
			i = i - j + 1
			j = 0
			continue
		}
		if j == needleLen {
			return i - needleLen
		}
	}
	return -1
}
