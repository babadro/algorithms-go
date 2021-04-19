package _1760_minimum_limit_of_balls_in_a_bag

// https://leetcode.com/problems/minimum-limit-of-balls-in-a-bag/discuss/1064702/JavaScript%3A-Binary-Search-solution-with-explanation
// passed. tptl. interesting usage of binary search
func minimumSize(nums []int, maxOperations int) int {
	l, r := 1, 0
	for _, num := range nums {
		if num > r {
			r = num
		}
	}

	for l < r {
		mid := l + (r-l)/2
		i, k := 0, 0
		for ; i < len(nums) && k <= maxOperations; i++ {
			k += (nums[i] - 1) / mid
		}

		if k <= maxOperations {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}
