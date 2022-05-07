package _744_find_smallest_letter_greater_than_target

import "testing"

func Test_nextGreatestLetter(t *testing.T) {
	tests := []struct {
		letters []byte
		target  byte
		want    byte
	}{
		{[]byte{'c', 'f', 'j'}, 'a', 'c'},
		{[]byte{'c', 'f', 'j'}, 'c', 'f'},
		{[]byte{'c', 'f', 'j'}, 'd', 'f'},
		{[]byte{'c', 'f', 'j'}, 'g', 'j'},
		{[]byte{'c', 'f', 'j'}, 'j', 'c'},
		{[]byte{'c', 'f', 'j'}, 'k', 'c'},
		{[]byte{'a', 'b'}, 'z', 'a'},
	}
	for _, tt := range tests {
		t.Run(string(tt.letters)+"_"+string(tt.target), func(t *testing.T) {
			if got := nextGreatestLetter(tt.letters, tt.target); got != tt.want {
				t.Errorf("nextGreatestLetter() = %s, want %s", string(got), string(tt.want))
			}
		})
	}
}
