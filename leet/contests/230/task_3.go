package _30

import "sort"

func minOperations(nums1 []int, nums2 []int) int {
	long, short := nums1, nums2
	if len(long) < len(short) {
		long, short = short, long
	}
	l, s := len(long), len(short)

	segments := l / 6
	if r := l % 6; r != 0 {
		segments++
	}

	if s < segments {
		return -1
	}

	sumLong, sumShort := 0, 0
	for _, num := range long {
		sumLong += num
	}
	for _, num := range short {
		sumShort += num
	}

	big, small := long, short
	sumBig, sumSmall := sumLong, sumShort
	if sumLong < sumShort {
		big, small = short, long
		sumBig, sumSmall = sumShort, sumLong
	}

	sort.Slice(big, func(i, j int) bool {
		return big[i] > big[j]
	})
	sort.Ints(small)
	bigL, smallL := len(big), len(small)

	counter, i, j := 0, 0, 0
	for sumBig != sumSmall {
		counter++
		diff := sumBig - sumSmall

		if i < bigL && j < smallL {
			down := big[i] - 1
			rise := 6 - small[j]
			if down > rise {
				if diff <= down {
					return counter
				}

				big[i] = 1
				i++
				sumBig -= down
			} else {
				if diff <= rise {
					return counter
				}

				small[j] = 6
				j++
				sumSmall += rise
			}
		} else if i < bigL {
			down := big[i] - 1
			if diff <= down {
				return counter
			}

			big[i] = 1
			i++
			sumBig -= down
		} else {
			rise := 6 - small[j]
			if diff <= rise {
				return counter
			}

			small[j] = 6
			j++
			sumSmall += rise
		}
	}

	return counter
}
