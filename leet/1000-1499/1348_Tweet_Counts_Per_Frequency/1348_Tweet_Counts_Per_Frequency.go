package _1348_Tweet_Counts_Per_Frequency

import "sort"

type tweet struct {
	times  []int
	sorted bool
}

type TweetCounts struct {
	m map[string]tweet
}

func Constructor() TweetCounts {

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
	}


}

func binarySearch(arr []int, target int) int {
	l, r := 0, len(arr)-1
	for l < r {
		m := l + (r-l)/2
		num := arr[m]
		if num
	}
}
