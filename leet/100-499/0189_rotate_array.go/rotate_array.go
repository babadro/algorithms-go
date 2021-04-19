package _189_rotate_array_go

// TODO 2 read and implement solution https://leetcode.com/problems/rotate-array/solution/

// Doesn't work
func rotate2(nums []int, k int) {
	l := len(nums)
	if l == 0 {
		return
	}
	k = k % l
	var idx int
	for i := 0; i < k; i++ {
		idx = l - k + i
		nums[i], nums[idx] = nums[idx], nums[i]
		for j := idx; j > k+i; j-- {
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}
	}
}

/*
input
[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53]
82

output
[25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,1,2,3,4,5]

expected
[25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24]
*/
