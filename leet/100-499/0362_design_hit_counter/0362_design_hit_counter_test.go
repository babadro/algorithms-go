package _362_design_hit_counter

import "testing"

/*
Input
["HitCounter", "hit", "hit", "hit", "getHits", "hit", "getHits", "getHits"]
[[], [1], [2], [3], [4], [300], [300], [301]]
Output
[null, null, null, null, 3, null, 4, 3]
*/
func TestHitCounter_Hit(t *testing.T) {
	c := Constructor()

	c.Hit(1)
	c.Hit(2)
	c.Hit(3)

	t.Log(c.GetHits(4)) // 3

	c.Hit(300)

	t.Log(c.GetHits(300)) // 4
	t.Log(c.GetHits(301)) // 3
}
