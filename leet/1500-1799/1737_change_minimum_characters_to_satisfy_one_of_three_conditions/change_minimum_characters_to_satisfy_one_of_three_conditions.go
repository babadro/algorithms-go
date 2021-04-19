package _1737_change_minimum_characters_to_satisfy_one_of_three_conditions

import "math"

//https://leetcode.com/contest/weekly-contest-225
// passed. fast enough
func minCharacters(a string, b string) int {
	res1, res2 := 0, 0

	mA, sumA, mB, sumB := [26]int{}, [27]int{}, [26]int{}, [27]int{}

	for _, char := range a {
		mA[char-'a']++
	}

	for _, char := range b {
		mB[char-'a']++
	}

	for i := 1; i < 27; i++ {
		sumA[i] = sumA[i-1] + mA[i-1]
		sumB[i] = sumB[i-1] + mB[i-1]
	}

	minCount, curr := math.MaxInt64, 0
	for i := 0; i <= 24; i++ {
		if curr = sum(sumA, i+1, 25) + sum(sumB, 0, i); curr < minCount {
			minCount = curr
		}

		if curr = sum(sumA, 0, i) + sum(sumB, i+1, 25); curr < minCount {
			minCount = curr
		}
	}

	res1 = minCount

	mostPopular := 0
	for i := 0; i < 26; i++ {
		count := mA[i] + mB[i]
		if count > mostPopular {
			mostPopular = count
		}
	}

	res2 = len(a) + len(b) - mostPopular

	if res2 < res1 {
		return res2
	}

	return res1
}

func sum(arr [27]int, from, to int) (res int) {
	return arr[to+1] - arr[from]
}
