package _880_check_if_word_equals_summation_of_two_words

import "strconv"

// passed. easy todo 3 try to find shorter solution
func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
	var num1 []byte
	for i := range firstWord {
		char := firstWord[i]
		num1 = append(num1, '0'+char-'a')
	}

	var num2 []byte
	for i := range secondWord {
		char := secondWord[i]
		num2 = append(num2, '0'+char-'a')
	}

	var num3 []byte
	for i := range targetWord {
		char := targetWord[i]
		num3 = append(num3, '0'+char-'a')
	}

	str1, str2, str3 := string(num1), string(num2), string(num3)

	n1, _ := strconv.Atoi(str1)
	n2, _ := strconv.Atoi(str2)
	n3, _ := strconv.Atoi(str3)

	return n1+n2 == n3
}
