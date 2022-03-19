package _2156_find_substring_with_given_hash_value

// tptl pass #hard #ntu
// https://leetcode.com/problems/find-substring-with-given-hash-value/discuss/1730321/JavaC%2B%2BPython-Sliding-Window-%2B-Rolling-Hash
func subStrHash(s string, power int, modulo int, k int, hashValue int) string {
	curHash, powerK := 0, 1
	res, n := 0, len(s)

	for i := n - 1; i >= 0; i-- {
		curHash = (curHash*power + val(s[i])) % modulo

		if i+k >= n {
			powerK = powerK * power % modulo
		} else { // #ntu why we add modulo here? Short answer: it prevents curHash - val(s[i+k])*powerK%modulo  from becoming negative.
			curHash = (curHash - val(s[i+k])*powerK%modulo + modulo) % modulo
		}

		if curHash == hashValue {
			res = i
		}
	}

	return s[res : res+k]
}

func val(char byte) int {
	return int(char-'a') + 1
}
