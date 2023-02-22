package __cyclic_sort

func cyclicSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for nums[i] != i+1 {
			tmp := nums[i]
			nums[i] = nums[nums[i]-1]
			nums[tmp-1] = tmp
		}
	}
}

func cyclicSort2(nums []int) {
	for i := 0; i < len(nums); {
		if j := nums[i] - 1; j != i {
			nums[i], nums[j] = nums[j], nums[i]
		} else {
			i++
		}
	}
}
