package _784_check_if_binary_string_has_at_most_one_segment_of_ones

func checkOnesSegment(s string) bool {
	i, n := 0, len(s)
	for ; i < n && s[i] == '0'; i++ {
	}

	for ; i < n && s[i] == '1'; i++ {
	}

	for ; i < n && s[i] == '0'; i++ {
	}

	return i == n
}
