package _014_Longest_Common_Prefix

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	end := 0
	var curr byte
Loop:
	for {
		curr = 0
		for i := range strs {
			if end == len(strs[i]) || (curr != 0 && strs[i][end] != curr) {
				break Loop
			}
			curr = strs[i][end]
		}
		end++
	}

	return strs[0][:end]
}
