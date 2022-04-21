package sorting

// stability and usefulness in parallel algorithms, distributed algorithms, and databases (ie. Merge join)
func merge(arr []int, l, m, r int) {
	n1 := m - l + 1
	n2 := r - m

	L, R := make([]int, n1), make([]int, n2)

	copy(L, arr[l:m+1])
	copy(R, arr[m+1:r+1])

	i, j := 0, 0
	k := l
	for i < n1 && j < n2 {
		if L[i] <= R[j] {
			arr[k] = L[i]
			i++
		} else {
			arr[k] = R[j]
			j++
		}
		k++
	}

	for i < n1 {
		arr[k] = L[i]
		i++
		k++
	}

	for j < n2 {
		arr[k] = R[j]
		j++
		k++
	}
}

func mergeSort(arr []int, l, r int) {
	if l < r {
		m := l + (r-l)/2

		mergeSort(arr, l, m)
		mergeSort(arr, m+1, r)

		merge(arr, l, m, r)
	}
}

func mergeSortParallel(arr []int, l, r, level, max int) {
	if l < r {
		m := l + (r-l)/2

		if level <= max {
			done := make(chan bool)

			go func() {
				mergeSortParallel(arr, l, m, level+1, max)
				done <- true
			}()

			mergeSortParallel(arr, m+1, r, level+1, max)
			<-done
		} else {
			mergeSort(arr, l, m)
			mergeSort(arr, m+1, r)
		}

		merge(arr, l, m, r)
	}
}
