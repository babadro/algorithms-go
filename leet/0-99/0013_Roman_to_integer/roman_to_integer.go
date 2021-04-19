package _013_Roman_to_integer

func romanToInt(s string) int {
	i := 0
	sum := 0
	for i < len(s) {
		if i != len(s)-1 {
			if s[i] == 'I' {
				if s[i+1] == 'V' {
					i += 2
					sum += 4
				} else if s[i+1] == 'X' {
					i += 2
					sum += 9
				} else {
					i++
					sum += 1
				}
				continue
			} else if s[i] == 'X' {
				if s[i+1] == 'L' {
					i += 2
					sum += 40
				} else if s[i+1] == 'C' {
					i += 2
					sum += 90
				} else {
					i++
					sum += 10
				}
				continue
			} else if s[i] == 'C' {
				if s[i+1] == 'D' {
					i += 2
					sum += 400
				} else if s[i+1] == 'M' {
					i += 2
					sum += 900
				} else {
					i++
					sum += 100
				}
				continue
			} else {
				sum += charToInt(s[i])
				i++
				continue
			}
		}
		sum += charToInt(s[i])
		i++
	}
	return sum
}

// Best solution
func romanToInt2(s string) int {

	tableRTI := map[rune]int{
		73: 1,
		86: 5,
		88: 10,
		76: 50,
		67: 100,
		68: 500,
		77: 1000,
	}
	sIntSlice := make([]int, len(s))

	for k, v := range s {
		sIntSlice[k] = tableRTI[v]
	}

	var output int
	var i int

	//The last elem of sIntSlice definitely is plus. So we don't need to judge on it.
	for ; i < len(s)-1; i++ {
		switch sIntSlice[i] {
		case 100, 10, 1:
			if sIntSlice[i] < sIntSlice[i+1] {
				output -= sIntSlice[i]
			} else {
				output += sIntSlice[i]
			}
		default:
			output += sIntSlice[i]
		}
	}
	output += sIntSlice[i]
	return output
}

func charToInt(b byte) int {
	switch b {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	}
	return 0
}
