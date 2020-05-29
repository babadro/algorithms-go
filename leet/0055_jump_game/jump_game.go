package _055_jump_game

// TODO 1 - doesn't work. infinitive loop
func canJump(nums []int) bool {
	i := 0
	last := len(nums) - 1
	for i < last {
		j := i + nums[i]
		if j >= last {
			return true
		}
		if nums[j] == 0 {
			for j = j - 1; j > i; j-- {
				if idx := j + nums[j]; idx >= last || nums[idx] != 0 {
					i = idx
					break
				}
			}
		} else {
			i = j
		}
	}
	return true
}
