package _299_bulls_and_cows

import "fmt"

func getHint(secret string, guess string) string {
	secretArr, guessArr := [10]int{}, [10]int{}
	A, B := 0, 0
	for i := range secret {
		secretChar, guessChar := secret[i], guess[i]
		if secretChar == guessChar {
			A++
			continue
		}
		secretArr[secretChar-'0']++
		guessArr[guessChar-'0']++
	}
	for i := 0; i < 10; i++ {
		B += min(secretArr[i], guessArr[i])
	}
	return fmt.Sprintf("%dA%dB", A, B)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
