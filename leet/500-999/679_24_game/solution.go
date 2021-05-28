package _679_24_game

const eps = 0.001

// passed. tptl. hard
func judgePoint24(nums []int) bool {
	list := make([]float64, 0, 4)
	for _, v := range nums {
		list = append(list, float64(v))
	}

	return search(list)
}

func search(nums []float64) bool {
	if len(nums) == 1 {
		return abs(nums[0]-24) < eps
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}

			var list []float64
			for k := 0; k < len(nums); k++ {
				if k != i && k != j {
					list = append(list, nums[k])
				}
			}

			if search(append(list, nums[i]+nums[j])) ||
				search(append(list, nums[i]-nums[j])) ||
				search(append(list, nums[i]*nums[j])) ||
				(nums[j] != 0 && search(append(list, nums[i]/nums[j]))) {
				return true
			}
		}
	}

	return false
}

func abs(a float64) float64 {
	if a > 0 {
		return a
	}

	return -a
}
