package _225

// todo 1 comp
//https://leetcode.com/contest/weekly-contest-225

func minCharacters(a string, b string) int {
	res1, res2, res3 := 0, 0, 0

	//	aTemp, bTemp := []byte(a), []byte(b)
	mA, mB := [26]int{}, [26]int{}
	sumA, sumB := [26][2]int{}, [26][2]int{}

	maxA, minA := byte('a'), byte('z')
	for i := range a {
		char := a[i]
		idx := char - 'a'
		mA[idx]++

		if a[i] < minA {
			minA = a[i]
		}
		if a[i] > maxA {
			maxA = a[i]
		}
	}

	sumA[1][0], sumB[1][0] = mA[1], mB[1]
	for i := 2; i < 26; i++ {
		sumA[i][0] += sumA[i-1][0]
		sumB[i][0] += sumB[i-1][0]
	}

	sumA[24][1], sumB[24][1] = mA[24], mB[24]
	for i := 23; i >= 0; i-- {
		sumA[i][1] += sumA[i+1][1]
		sumB[i][1] += sumB[i+1][1]
	}

	maxB, minB := byte('a'), byte('z')
	for i := range b {
		char := b[i]
		idx := char - 'a'
		mB[idx]++
		if b[i] < minB {
			minB = b[i]
		}
		if b[i] > maxB {
			maxB = b[i]
		}
	}

	for i, j := 0, 25; i <= j; {
		countA, countB := mA[i], mB[j]
		if countA < countB {
			res1 += countA
			i++
		} else if countB > countA {
			res1 += countB
			j--
		} else {
			if i == 25 {
				res1 += countA
				i++
			} else if j == 0 {
				res1 += countB
				j--
			} else if mA[i+1] < mB[j] {
				res1 += countA
				i++
			} else {
				res1 += countB
				j--
			}
		}
	}

	for i, j := 0, 25; i <= j; {
		countA, countB := mB[i], mA[j]
		if countA < countB {
			res2 += countA
			i++
		} else if countB > countA {
			res2 += countB
			j--
		} else {
			if i == 25 {
				res2 += countA
				i++
			} else if j == 0 {
				res2 += countB
				j--
			} else if mA[i+1] < mB[j] {
				res2 += countA
				i++
			} else {
				res2 += countB
				j--
			}
		}
	}

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

	res3 = len(concat) - mostPopular

	res := res1
	if res2 < res {
		res = res2
	}
	if res3 < res {
		res = res3
	}

	return res
}
