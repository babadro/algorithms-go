package _2116_check_if_a_parentheses_string_can_be_valid

// #bnsrg passed
// todo2 check right-to-left and left-to-righ solution:
// https://leetcode.com/problems/check-if-a-parentheses-string-can-be-valid/solutions/1646594/left-to-right-and-right-to-left/
func canBeValid(s string, lock string) bool {
	var openParentheses, unlocked []int

	for i := range s {
		switch {
		case lock[i] == '0':
			unlocked = append(unlocked, i)
		case s[i] == '(':
			openParentheses = append(openParentheses, i)
		case len(openParentheses) > 0:
			openParentheses = openParentheses[:len(openParentheses)-1]
		case len(unlocked) > 0:
			unlocked = unlocked[:len(unlocked)-1]
		default:
			return false
		}
	}

	for len(openParentheses) > 0 {
		lastUnlocked, lastOpenParentheses := len(unlocked)-1, len(openParentheses)-1
		if len(unlocked) == 0 || openParentheses[lastOpenParentheses] > unlocked[lastUnlocked] {
			return false
		}

		unlocked = unlocked[:lastUnlocked]
		openParentheses = openParentheses[:lastOpenParentheses]
	}

	return len(openParentheses) == 0 && len(unlocked)%2 == 0
}
