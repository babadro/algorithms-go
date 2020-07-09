package _078_subsets

// 100%, 27%
// TODO 3 look at another approaches in solution
func subsets(nums []int) [][]int {
	output := [][]int{{}}

	for _, num := range nums {
		for _, curr := range output {
			output = append(output, append([]int{num}, curr...))
		}
	}

	return output
}
