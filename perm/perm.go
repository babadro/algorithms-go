package perm

func NextPermutationBytes(arr []byte) (ok bool) {
	n := len(arr)
	i := n - 2
	for ; i >= 0 && arr[i] >= arr[i+1]; i-- {
	}

	if i >= 0 {
		ok = true

		j := n - 1
		for ; j >= 0 && arr[j] <= arr[i]; j-- {
		}

		arr[i], arr[j] = arr[j], arr[i]
	}

	for k, m := i+1, n-1; k < m; k, m = k+1, m-1 {
		arr[k], arr[m] = arr[m], arr[k]
	}

	return ok
}
