package _941_valid_mountain_array

// tptl. passed
func validMountainArray(arr []int) bool {
	n := len(arr)
	if n < 3 {
		return false
	}

	up, down := false, false
	for i := 1; i < n; i++ {
		switch {
		case arr[i] == arr[i-1]:
			return false
		case !down && arr[i] < arr[i-1]:
			down = true
		case !up && arr[i] > arr[i-1]:
			up = true
		}

		if down {
			if !up {
				return false
			}

			if arr[i] > arr[i-1] {
				return false
			}
		} else if arr[i] < arr[i-1] {
			return false
		}
	}

	return down
}
