package _165_compare_version_numbers

import "testing"

/*
Example 1:

Input: version1 = "0.1", version2 = "1.1"
Output: -1
Example 2:

Input: version1 = "1.0.1", version2 = "1"
Output: 1
Example 3:

Input: version1 = "7.5.2.4", version2 = "7.5.3"
Output: -1
Example 4:

Input: version1 = "1.01", version2 = "1.001"
Output: 0
Explanation: Ignoring leading zeroes, both “01” and “001" represent the same number “1”
Example 5:

Input: version1 = "1.0", version2 = "1.0.0"
Output: 0
Explanation: The first version number does not have a third level revision number, which means its third level revision number is default to "0"
*/

func TestCompareVersion(t *testing.T) {
	cases := []struct {
		ver1     string
		ver2     string
		expected int
	}{
		//{"0.1", "1.1", 0},
		{"1.0.1", "1", 1},
		//{"7.5.2.4", "7.5.3", -1},
		//{"1.01", "1.001", 0},
		//{"1.0", "1.0.0", 0},
		//{"000000", "00.00.00.00", 0},
		//{"1001.10.00", "1001.1.01", 1},
		//{"1.0.1", "1.000.1", 0},
		//{"1.0", "1.000", 0},
	}

	for i, c := range cases {
		if fact := CompareVersion(c.ver1, c.ver2); fact != c.expected {
			t.Errorf("test#%d error: want %d, got %d, ver1: %s, ver2: %s", i+1, c.expected, fact, c.ver1, c.ver2)
		}
	}
}
