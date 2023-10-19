package _2842_count_k_subsequences_of_a_string_with_maximum_beauty

import "sort"

const mod = 1_000_000_007

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

	sort.Slice(freqArr, func(i, j int) bool {
		return freqArr[i] > freqArr[j]
	})

	res := 1
	lastResBeforeFreqChange := 1
	lastIDxBeforeFreqChange := -1

	i := 0
	for ; i < subSeqLen; i++ {
		if i > 0 && freqArr[i] != freqArr[i-1] {
			lastResBeforeFreqChange = res
			lastIDxBeforeFreqChange = i - 1
		}

		res = res * freqArr[i] % mod
	}

	if i == len(freqArr) || freqArr[i] != freqArr[i-1] {
		return res
	}

	k := i - lastIDxBeforeFreqChange - 1

	freq := freqArr[i]

	for i < len(freqArr) && freqArr[i] == freqArr[i-1] {
		i++
	}

	n := i - lastIDxBeforeFreqChange - 1

	if lastIDxBeforeFreqChange != -1 {
		//res = lastResBeforeFreqChange * k % mod
		res = lastResBeforeFreqChange * freq % mod
	}

	// n!/((n-k)!*k!)
	numberOfCombinations := combinationNOverKMod(n, k)

	return res * numberOfCombinations % mod
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
