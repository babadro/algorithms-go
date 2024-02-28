package _2842_count_k_subsequences_of_a_string_with_maximum_beauty

import "sort"

const mod = 1_000_000_007

// #bnsrg passed
// todo 2 study solution.go file - it doesn't use complicated formula
func countKSubsequencesWithMaxBeauty(s string, subSeqLen int) int {
	freqMap := make(map[byte]int)
	for i := range s {
		freqMap[s[i]]++
	}

	if len(freqMap) < subSeqLen {
		return 0
	}

	freqArr := make([]int, 0, len(freqMap))
	for _, freq := range freqMap {
		freqArr = append(freqArr, freq)
	}

	freqFreq := make(map[int]int)
	for _, freq := range freqArr {
		freqFreq[freq]++
	}

	sort.Slice(freqArr, func(i, j int) bool {
		return freqArr[i] > freqArr[j]
	})

	minFreq := freqArr[subSeqLen-1]

	minFreqCount := freqFreq[minFreq]

	res := 1
	for i := 0; i < subSeqLen; i++ {
		freq := freqArr[i]

		if freq != minFreq {
			res = res * freq % mod

			continue
		}

		remainingFreqs := subSeqLen - i

		for j := 0; j < remainingFreqs; j++ {
			res = res * freq % mod
		}

		c := combinationNOverKMod(minFreqCount, remainingFreqs)

		res = res * c % mod

		return res
	}

	return res
}

func powerMod(x, y int) int {
	res := 1
	base := x % mod

	for y > 0 {
		if y%2 == 1 {
			res = res * base % mod
		}

		base = base * base % mod

		y /= 2
	}

	return res
}

func factorialMod(n int) int {
	res := 1

	for i := 2; i <= n; i++ {
		res = res * i % mod
	}

	return res
}

func combinationNOverKMod(n, k int) int {
	if k > n {
		return 0
	}

	num := factorialMod(n)

	den := factorialMod(k) * factorialMod(n-k) % mod

	return num * powerMod(den, mod-2) % mod
}
