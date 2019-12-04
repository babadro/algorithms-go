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

		flag1, flag2 := false, false
		for {
			if i < len(version1) && version1[i] != 46 {
				i++
			} else if !flag1 && version1[i-1] > 0 {
				flag1 = true
			}
			if j < len(version2) && version2[j] != 46 {
				j++
			} else if !flag2 && version2[j-1] > 0 {
				flag2 = true
			}

			if (i == len(version1) && j == len(version2)) || (version1[i] == 46 && version2[j] == 46) {
				i++
				j++
				break
			}
			if (!flag1 && flag2) || (version1[i] > version2[j]) {
				return 1
			} else if (flag1 && !flag2) || (version1[i] < version2[j]) {
				return -1
			}
		}

	}
}
