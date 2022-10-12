package _1348_Tweet_Counts_Per_Frequency

import "sort"

type tweet struct {
	times  []int
	sorted bool
}

type TweetCounts struct {
	m map[string]tweet
}

// tptl. passed, but slow
// todo 2 find faster (and simpler) solution
func Constructor() TweetCounts {
	return TweetCounts{m: make(map[string]tweet)}
}

func (t *TweetCounts) RecordTweet(tweetName string, time int) {
	tw := t.m[tweetName]
	tw.times = append(tw.times, time)
	tw.sorted = false

	t.m[tweetName] = tw
}

func (t *TweetCounts) GetTweetCountsPerFrequency(freq, tweetName string, startTime, endTime int) []int {
	tw := t.m[tweetName]
	if !tw.sorted {
		sort.Ints(tw.times)
		tw.sorted = true
		t.m[tweetName] = tw
	}

	var idx int
	tweetTime := -1
	if len(tw.times) > 0 {
		idx = binarySearchStart(tw.times, startTime)
		tweetTime = tw.times[idx]
	}

	ch := chunk(freq)
	var res []int
	for beginChunkTime := startTime; beginChunkTime <= endTime; beginChunkTime += ch {
		next := beginChunkTime + ch
		if tweetTime >= next || tweetTime < beginChunkTime {
			res = append(res, 0)
			continue
		}

		counter := 0
		for ; idx < len(tw.times); idx++ {
			if tweetTime = tw.times[idx]; tweetTime >= next || tweetTime > endTime {
				break
			}

			counter++
		}

		res = append(res, counter)
	}

	return res
}

func binarySearchStart(arr []int, target int) int {
	l, r := 0, len(arr)-1
	for l < r {
		m := l + (r-l)/2
		num := arr[m]
		if num >= target {
			r = m
		} else {
			l = m + 1
		}
	}

	return l
}

func chunk(str string) int {
	switch str {
	case "minute":
		return 60
	case "hour":
		return 3600
	case "day":
		return 86400
	default:
		return -1
	}
}
