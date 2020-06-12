package _482_license_key_formatting

// 4 ms - 82.8%, 4.3MB - 97.6%
func licenseKeyFormatting(S string, K int) string {
	b := make([]byte, len(S)+len(S)/K)
	counter, idx := 0, len(b)-1
	for i := len(S) - 1; i >= 0; i-- {
		char := S[i]
		if char == '-' {
			continue
		}
		if char >= 'a' && char <= 'z' {
			char -= 32
		}
		if counter%K == 0 && counter > 0 {
			b[idx] = '-'
			idx--
		}
		b[idx] = char
		idx--
		counter++
	}
	return string(b[idx+1:])
}
