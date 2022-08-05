package _911_online_election

type TopVotedCandidate struct {
	totalToLastVotes [][]vote
}

type vote struct {
	person, time int
}

// todo 1 doesn't work fix bug (second testcase failed)
func Constructor(persons []int, times []int) TopVotedCandidate {
	totalToLastVotes := make([][]vote, 1)
	count := make(map[int]int)
	for i := range persons {
		v := vote{
			person: persons[i],
			time:   times[i],
		}

		count[v.person]++

		c := count[v.person]

		if len(totalToLastVotes) == c {
			totalToLastVotes = append(totalToLastVotes, make([]vote, 0))
		}

		totalToLastVotes[c] = append(totalToLastVotes[c], v)
	}

	return TopVotedCandidate{totalToLastVotes: totalToLastVotes}
}

func (this *TopVotedCandidate) Q(t int) int {
	l, r := 1, len(this.totalToLastVotes)-1
	for l < r {
		m := l + (r-l)/2
		if this.totalToLastVotes[m][0].time > t {
			r = m
		} else {
			l = m + 1
		}
	}

	votes := this.totalToLastVotes[l-1]

	l, r = 0, len(votes)-1
	for l < r {
		m := l + (r-l)/2
		if votes[m].time > t {
			r = m
		} else {
			l = m + 1
		}
	}

	return votes[l-1].person
}
