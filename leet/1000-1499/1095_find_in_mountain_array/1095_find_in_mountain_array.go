package _1095_find_in_mountain_array

type MountainArray struct {
	arr []int
}

func (this MountainArray) get(index int) int {
	return this.arr[index]
}
func (this MountainArray) length() int {
	return len(this.arr)
}

func findInMountainArray(target int, arr *MountainArray) int {
	l, r := 0, arr.length()-1
	for l < r {
		m := (l + r) / 2
		if arr.get(m) > arr.get(m+1) {
			r = m
		} else {
			l = m + 1
		}
	}

	peak := l
	idx := binarySearch(target, 0, peak, true, arr)
	if idx != -1 {
		return idx
	}

	return binarySearch(target, peak, arr.length()-1, false, arr)
}

func binarySearch(target int, l, r int, asc bool, arr *MountainArray) int {
	var diff int
	for l <= r {
		m := l + (r-l)/2
		num := arr.get(m)
		if asc {
			diff = num - target
		} else {
			diff = target - num
		}

		if diff > 0 {
			r = m - 1
		} else if diff < 0 {
			l = m + 1
		} else {
			return m
		}
	}

	return -1
}
