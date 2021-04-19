package _836_rectangle_overlap

import "github.com/babadro/algorithms-go/utils"

// tptl. passed.
func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	return utils.Min(rec1[2], rec2[2]) > utils.Max(rec1[0], rec2[0]) &&
		utils.Min(rec1[3], rec2[3]) > utils.Max(rec1[1], rec2[1])
}
