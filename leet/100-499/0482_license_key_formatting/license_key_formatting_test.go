package _482_license_key_formatting

import "testing"

func TestLicenseKeyFormatting(t *testing.T) {
	cases := []struct {
		input, expected string
		K               int
	}{
		{"1243", "1243", 4},
		{"asdf", "A-S-D-F", 1},
		{"5F3Z-2e-9-w", "5F3Z-2E9W", 4},
		{"2-5g-3-J", "2-5G-3J", 2},
		{"2-5g-3-J", "25G3J", 5},
		{"sdf3-a-aFD-34ff-2Fd-3fd-dfDF", "S-DF3-AAF-D34-FF2-FD3-FDD-FDF", 3},
		{"AS-d-f", "A-SDF", 3},
		{"A", "A", 3},
		{"ab", "AB", 3},
		{"abcd", "A-BCD", 3},
	}

	for i, c := range cases {
		if fact := licenseKeyFormatting(c.input, c.K); fact != c.expected {
			t.Errorf("case#%d, want %s, got %s", i+1, c.expected, fact)
		}
	}
}
