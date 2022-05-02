package __Introduction

func findMissingNumber(arr []int) int {
	n := len(arr) + 1
	res := 1
	for i := 2; i <= n; i++ {
		res ^= i
	}

	for i := range arr {
		res ^= arr[i]
	}

	return res
}
