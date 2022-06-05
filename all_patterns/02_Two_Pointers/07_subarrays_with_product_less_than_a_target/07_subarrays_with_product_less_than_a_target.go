package _7_subarrays_with_product_less_than_a_target

// see leetcode 713
func findSubarrays(arr []int, target int) [][]int {
	left, res, prod := 0, make([][]int, 0), 1
	for right := range arr {
		prod *= arr[right]
		for prod >= target && left <= right {
			prod /= arr[left]
			left++
		}

		if left <= right {
			length := right - left + 1
			tmp := make([]int, length)
			for j := right; j >= left; j-- {
				tmp[j-left] = arr[j]
				newArr := make([]int, right-j+1)
				copy(newArr, tmp[j-left:])
				res = append(res, newArr)
			}
		}
	}

	return res
}
