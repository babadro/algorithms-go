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
		for version1[i] == 48 && i < len(version1) {
			i++
		}
		for version2[j] == 48 && j < len(version2) {
			j++
		}
		for {
			if version1[i] < len(version1) && ver == 46
		}

	}
}
