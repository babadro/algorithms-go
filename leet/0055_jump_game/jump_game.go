package _055_jump_game

const (
	Unknown = iota
	Good
	Bad
)

var memo []int

// Approach 3: Dynamic Programming Bottom-up
// 5.17%, 14.29%
func canJumpMemo(nums []int) bool {
	memo = make([]int, len(nums))
	memo[len(nums)-1] = Good
	return canJumpFromPositionMemo(0, nums)
}

func canJumpFromPositionMemo(position int, nums []int) bool {
	if v := memo[position]; v == Good {
		return true
	} else if v == Bad {
		return false
	}

	furthestJump := position + nums[position]
	if furthestJump >= len(nums)-1 {
		return true
	}
	for nextPosition := position + 1; nextPosition <= furthestJump; nextPosition++ {
		if canJumpFromPositionMemo(nextPosition, nums) {
			memo[position] = Good
			return true
		}
	}

	memo[position] = Bad
	return false
}

// Approach 1: Backtracking. Time limit exceeded
func canJumpBacktracing(nums []int) bool {
	return canJumpFromPositionBacktracing(0, nums)
}

func canJumpFromPositionBacktracing(position int, nums []int) bool {
	if position == len(nums)-1 {
		return true
	}

	furthestJump := position + nums[position]
	if furthestJump >= len(nums)-1 {
		return true
	}
	for nextPosition := furthestJump; nextPosition > position; nextPosition-- {
		if canJumpFromPositionBacktracing(nextPosition, nums) {
			return true
		}
	}

	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// TODO 1 - doesn't work. infinitive loop
func canJumpOld(nums []int) bool {
	i := 0
	last := len(nums) - 1
Loop:
	for i < last {
		j := i + nums[i]
		if j >= last {
			return true
		}
		if nums[j] == 0 {
			for j = j - 1; j > i; j-- {
				if idx := j + nums[j]; idx >= last || nums[idx] != 0 {
					i = idx
					continue Loop
				}
			}
			return false
		} else {
			i = j
		}
	}
	return true
}
