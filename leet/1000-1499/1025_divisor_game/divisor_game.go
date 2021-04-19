package _1025_divisor_game

// passed. tptl
// A simple explanation would be that, we can't go from an odd number to an odd number but we can always go from an even number to an even number.
// So if Alice starts from an even number, she will always try to stay at an even number in her turn and she starts at an odd number, she will always remain at an odd number.
// 1 is an odd number and if anyone reaches 1, they loose. So basically if you start at an odd number you loose since the other person will always opt to be at an even number until you reach one and loose.
func divisorGame(N int) bool {
	return N%2 == 0
}
