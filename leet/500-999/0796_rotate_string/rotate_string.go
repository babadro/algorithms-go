package _796_rotate_string

import (
	"bytes"
	"strings"
)

// TODO 2 read solutions - some interesting algs

// 100%, 19.67%
func rotateString(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}
	return strings.Contains(A+A, B)
}

// brute force 100%, 76%
func rotateStringBruteForce(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}
	if len(A) == 0 {
		return true
	}
	a1, b1 := []byte(A), []byte(B)
	for i := 0; i < len(B); i++ {
		a1 = append(a1, a1[0])
		a1 = a1[1:]
		if bytes.Equal(a1, b1) {
			return true
		}
	}
	return false
}
