package _39

import (
	"github.com/babadro/algorithms-go/utils"
	"sort"
)

// passed. todo 2 check better solution
func getMinDistance(nums []int, target int, start int) int {
	var idxes []int
	for i, num := range nums {
		if num == target {
			idxes = append(idxes, i)
		}
	}

	sort.Slice(idxes, func(i, j int) bool {
		return utils.Abs(idxes[i]-start) < utils.Abs(idxes[j]-start)
	})

	return utils.Abs(idxes[0] - start)
}
