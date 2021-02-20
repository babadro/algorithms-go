package _1712_ways_to_split_array_into_three_subarrays

func waysToSplit(nums []int) int {
	n := len(nums)
	for i := 1; i < n; i++ {
		nums[i] += nums[i-1]
	}

	counter := 0
	for i := 0; i < n-2; i++ {
		start, finish := -1, n-1
		for j := i + 1; j < n-1; j++ {
			if start == -1 {
				if nums[j]-nums[i] >= nums[i] {
					start = j
				}
			} else if nums[j]-nums[i] > nums[n-1]-nums[j] {
				finish = j
				break
			}
		}

		if start == -1 || nums[n-1]-nums[finish-1] < nums[finish-1]-nums[i] {
			continue
		}

		counter += finish - start
	}

	return counter % 1_000_000_007
}

//if nums[i] >= nums[n-1]-nums[i] {
//	break
//}
//tmp := nums[i+1:n-2]
//nTmp := len(tmp)
//
//start := sort.Search(nTmp, func(j int) bool {
//	return tmp[j]-nums[i] >= nums[i]
//})
//
//finish := sort.Search(nTmp, func(j int) bool {
//	return tmp[j]-nums[i] > nums[n-1]-tmp[j]
//})
