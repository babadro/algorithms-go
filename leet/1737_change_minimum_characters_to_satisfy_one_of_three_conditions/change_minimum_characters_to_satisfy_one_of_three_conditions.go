package _1737_change_minimum_characters_to_satisfy_one_of_three_conditions

import "math"

//https://leetcode.com/contest/weekly-contest-225
// passed but not fast.
func minCharacters(a string, b string) int {
	res1, res2 := 0, 0

	mA, mB := [26]int{}, [26]int{}

	for i := range a {
		char := a[i]
		idx := char - 'a'
		mA[idx]++
	}

	for i := range b {
		char := b[i]
		idx := char - 'a'
		mB[idx]++
	}

	min, curr := math.MaxInt64, 0
	for i := 0; i <= 24; i++ {
		if curr = sum(mA, i+1, 25) + sum(mB, 0, i); curr < min {
			min = curr
		}

		if curr = sum(mA, 0, i) + sum(mB, i+1, 25); curr < min {
			min = curr
		}
	}

	res1 = min

	concat := a + b
	counter := make(map[byte]int)
	for i := range concat {
		char := concat[i]
		counter[char]++
	}

	mostPopular := 0
	for _, v := range counter {
		if v > mostPopular {
			mostPopular = v
		}
	}

	res2 = len(concat) - mostPopular

	if res2 < res1 {
		return res2
	}

	return res1
}

func sum(arr [26]int, from, to int) (res int) {
	for i := from; i <= to; i++ {
		res += arr[i]
	}

	return res
}
