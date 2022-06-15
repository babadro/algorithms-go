package _015_3Sum

import "sort"

// tptl
func threeSum(arr []int) [][]int {
	sort.Ints(arr)
	var triplets [][]int
	for i := 0; i < len(arr)-2; i++ {
		if i > 0 && arr[i] == arr[i-1] { // skip duplicates
			continue
		}

		triplets = searchPair(arr, -arr[i], i+1, triplets)
	}

	return triplets
}

func searchPair(arr []int, targetSum, left int, triplets [][]int) [][]int {
	right := len(arr) - 1
	for left < right {
		currSum := arr[left] + arr[right]
		if currSum == targetSum {
			triplets = append(triplets, []int{-targetSum, arr[left], arr[right]})
			left++
			right--

			for left < right && arr[left] == arr[left-1] {
				left++
			}

			for left < right && arr[right] == arr[right+1] {
				right--
			}
		} else if targetSum > currSum {
			left++
		} else {
			right--
		}
	}

	return triplets
}
