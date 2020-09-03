package _179_largest_number

import (
	"sort"
	"strconv"
	"strings"
)

// Shortest but not fast: 60%, 69%
// TODO 2 look for solutions with better performance
func largestNumber(nums []int) string {
	strNums := make([]string, len(nums))
	for i, num := range nums {
		strNums[i] = strconv.Itoa(num)
	}
	sort.Slice(strNums, func(i, j int) bool {
		return strNums[i]+strNums[j] > strNums[j]+strNums[i]
	})
	if strNums[0] == "0" {
		return "0"
	}

	return strings.Join(strNums, "")
}
