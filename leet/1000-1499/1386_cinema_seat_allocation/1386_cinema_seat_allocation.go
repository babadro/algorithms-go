package _1386_cinema_seat_allocation

var m = [16]byte{
	0:  2,
	1:  1,
	2:  1,
	3:  1,
	4:  1,
	5:  0,
	6:  0,
	7:  0,
	8:  1,
	9:  1,
	10: 0,
	11: 0,
	12: 1,
	13: 0,
	14: 0,
	15: 0,
}

// #bnsrg medium passed
// todo2 there are shorter solutions - need to check them
func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
	rowMasks := make(map[int]byte)
	for _, r := range reservedSeats {
		rowN, seatN := r[0], r[1]

		bit := byte(0)

		switch seatN {
		case 8, 9:
			bit = 1
		case 6, 7:
			bit = 2
		case 4, 5:
			bit = 4
		case 2, 3:
			bit = 8
		}

		val := rowMasks[rowN]

		val |= bit

		rowMasks[rowN] = val
	}

	res := (n - len(rowMasks)) * 2

	for _, mask := range rowMasks {
		res += int(m[mask])
	}

	return res
}
