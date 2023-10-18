package _2842_count_k_subsequences_of_a_string_with_maximum_beauty

import "sort"

const mod = 1_000_000_007

func countKSubsequencesWithMaxBeauty(s string, subSeqLen int) int {
	freqMap := make(map[byte]int)
	for i := range s {
		freqMap[s[i]]++
	}

	freqArr := make([]int, 0, len(freqMap))
	for _, freq := range freqMap {
		freqArr = append(freqArr, freq)
	}

	sort.Slice(freqArr, func(i, j int) bool {
		return freqArr[i] > freqArr[j]
	})

	res := 1
	lastResBeforeFreqChange := 0
	lastIDxBeforeFreqChange := 0

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

	k := i - lastIDxBeforeFreqChange + 1

	freq := freqArr[i]

	for i < len(freqArr) && freqArr[i] == freqArr[i-1] {
		i++
	}

	n := i - lastIDxBeforeFreqChange + 1

	res = lastResBeforeFreqChange * k % mod
	res = res * freq % mod

	// n!/((n-k)!*k!)
	numberOfCombinationsMod := ? // todo write the func to compute nCk % MOD


	return res * numberOfCombinations % mod
}