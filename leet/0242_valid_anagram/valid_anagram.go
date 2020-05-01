package _242_valid_anagram

import (
	"sort"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	arr := [26]int{}
	for i := 0; i < len(s); i++ {
		arr[s[i]-'a']++
		arr[t[i]-'a']--
	}
	for _, count := range arr {
		if count != 0 {
			return false
		}
	}
	return true
}

func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	b1 := []byte(s)
	b2 := []byte(t)
	sort.Slice(b1, func(i, j int) bool {
		return b1[i] < b1[j]
	})
	sort.Slice(b2, func(i, j int) bool {
		return b2[i] < b2[j]
	})
	for i := 0; i < len(s); i++ {
		if b1[i] != b2[i] {
			return false
		}
	}
	return true
}

func isAnagramNaive(s string, t string) bool {
	len1, len2 := len(s), len(t)
	if len1 != len2 {
		return false
	}
	map1, map2 := make(map[int32]int, len1), make(map[int32]int, len2)
	for _, b := range s {
		map1[b]++
	}
	for _, b := range t {
		map2[b]++
	}
	if len(map1) != len(map2) {
		return false
	}
	for k, v := range map1 {
		if map2[k] != v {
			return false
		}
	}
	return true
}
