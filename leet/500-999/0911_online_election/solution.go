package _911_online_election

type TopVotedCandidate struct {
	A [][2]int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	var A [][2]int
	count := make(map[int]int)
	for i := range persons {
		person, time := persons[i], times[i]

		c := count[person]
		count[person]++

		for len(A) <= c {
			A = append(A, [][2]int{}...)
		}

		A[c] = [2]int{person, time}
	}
}

func (this *TopVotedCandidate) Q(t int) int {

}
