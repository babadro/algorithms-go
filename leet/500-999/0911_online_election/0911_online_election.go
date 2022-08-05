package _911_online_election

type TopVotedCandidate2 struct {
	votes map[int][]int
}

// passed, but slow
func Constructor2(persons []int, times []int) TopVotedCandidate2 {
	votes := make(map[int][]int)
	for i := range persons {
		person := persons[i]
		votes[person] = append(votes[person], times[i])
	}

	return TopVotedCandidate2{votes: votes}
}

func (this *TopVotedCandidate2) Q(t int) int {
	maxVotes := 0
	var winners [][2]int
	for person, votes := range this.votes {
		lastIDx := binarySearch(votes, t)
		votesCount := lastIDx + 1
		if votesCount == 0 {
			continue
		}

		if votesCount > maxVotes {
			maxVotes = votesCount
			winners = winners[:0]
			winners = append(winners, [2]int{person, votes[lastIDx]})
		} else if votesCount == maxVotes {
			winners = append(winners, [2]int{person, votes[lastIDx]})
		}
	}

	winner, lastVoteTime := -1, -1
	for _, w := range winners {
		if w[1] > lastVoteTime {
			winner, lastVoteTime = w[0], w[1]
		}
	}

	return winner
}

func binarySearch(votes []int, t int) int {
	l, r := 0, len(votes)-1
	for l <= r {
		m := l + (r-l)/2
		if votes[m] > t {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return r
}
