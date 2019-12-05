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
	segment1, segment2 := 0, 0
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
		flag1, flag2 := false, false
		for {
			if i < len(version1) && version1[i] != 46 {
				val1 = version1[i]
				i++

			} else if version1[i-1] > 48 {
				flag1 = true
			}
			if j < len(version2) && version2[j] != 46 {
				val2 = version2[j]
				j++

			} else if version2[j-1] > 48 {
				flag2 = true
			}

			if (!flag1 && flag2) || (val1 > val2) {
				return 1
			} else if (flag1 && !flag2) || (val1 < val2) {
				return -1
			}
			if (i == len(version1) || version1[i] == 46) && (j == len(version2) || version2[j] == 46) {
				if i < len(version1) {
					i++
					segment1++
				}
				if j < len(version2) {
					j++
					segment2++
				}
				break
			}
		}

	}
}
