package _165_compare_version_numbers

/*
	{"0.1", "1.1", 0},
	{"1.0.1", "1", 1},
	{"7.5.2.4", "7.5.3", -1},
	{"1.01", "1.001", 0},
	{"1.0", "1.0.0", 0},
	{"000000", "00.00.00.00"}
*/

// "0" = 48, "." = 46
func CompareVersion(version1 string, version2 string) int {
	i, j := 0, 0
	for {
		for i < len(version1) && version1[i] == 48 {
			i++
		}
		for j < len(version2) && version2[j] == 48 {
			j++
		}
		if i == len(version1) && j == len(version2) {
			return 0
		}

		val1, val2 := uint8(0), uint8(0)

		for {
			if i < len(version1) && version1[i] != 46 {
				val1 *= 10
				val1 += version1[i] - 48
				i++
			}
			if j < len(version2) && version2[j] != 46 {
				val2 *= 10
				val2 += version2[j] - 48
				j++

			}

			if (i == len(version1) || version1[i] == 46) && (j == len(version2) || version2[j] == 46) {
				if val1 > val2 {
					return 1
				}
				if val1 < val2 {
					return -1
				}
				if i < len(version1) {
					i++
				}
				if j < len(version2) {
					j++
				}
				break
			}
		}
	}
}

func CompareVersion2(version1 string, version2 string) int {
	i, j := 0, 0
	for i < len(version1) || j < len(version2) {
		val1, val2 := uint8(0), uint8(0)
		for i < len(version1) && version1[i] != 46 {
			val1 *= 10
			val1 += version1[i] - 48
			i++
		}
		for j < len(version2) && version2[j] != 46 {
			val2 *= 10
			val2 += version2[j] - 48
			j++

		}

		if val1 > val2 {
			return 1
		}
		if val1 < val2 {
			return -1
		}
		i++
		j++
	}
	return 0

}
