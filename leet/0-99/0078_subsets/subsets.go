package _078_subsets

// tptl passed best solution
func subsets(nums []int) [][]int {
	res := [][]int{{}}
	for _, num := range nums {
		size := len(res)
		for i := 0; i < size; i++ {
			set := make([]int, len(res[i])+1)
			copy(set, res[i])
			set[len(set)-1] = num

			res = append(res, set)
		}
	}

	return res
}

// 100%, 27% not obvious solution
// TODO 3 look at another approaches in solution
func subsets3(nums []int) [][]int {
	output := [][]int{{}}

	for _, num := range nums {
		for _, curr := range output {
			output = append(output, append([]int{num}, curr...))
		}
	}

	return output
}
