package _062_unique_paths

import "testing"

func TestPerm(t *testing.T) {
	arr := []byte{1, 2, 3, 1, 0, 1}
	var res [][]byte
	distinctPermutations(arr, &res, 0, len(arr)-1)
	t.Log(res)
	count := 0
	permCount(arr, &count, 0, len(arr)-1)
	t.Log(count)
}

func TestUniquePaths(t *testing.T) {
	t.Log(uniquePathsByPerm(3, 1))
}
