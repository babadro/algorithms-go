package _1348_Tweet_Counts_Per_Frequency

import (
	"math"
	"sort"
)

const (
	minute = "minute"
	hour   = "hour"
	day    = "day"
)

type TweetCounts struct {
	tweetMap map[string][]int
}

func Constructor() TweetCounts {
	return TweetCounts{tweetMap: make(map[string][]int)}
}

func (this *TweetCounts) RecordTweet(tweetName string, time int) {
	this.tweetMap[tweetName] = append(this.tweetMap[tweetName], time)
	sort.Slice(this.tweetMap[tweetName], func(i, j int) bool {
		return this.tweetMap[tweetName][i] < this.tweetMap[tweetName][j]
	})
}

func (this *TweetCounts) GetTweetCountsPerFrequency(freq string, tweetName string, startTime int, endTime int) []int {
	intervals := Intervals(freq, startTime, endTime)
	tweetTimes := this.tweetMap[tweetName]
	var res []int
	for _, interval := range intervals {
		count := 0
		for _, t := range tweetTimes {
			if t >= interval.start && t < interval.end {
				count++
			}
		}
		res = append(res, count)
	}
	return res
}

type Interval struct {
	start, end int
}

func Intervals(freq string, start, end int) []Interval {
	end++
	diff := end - start
	if diff <= 0 {
		return nil
	}
	var freqInSeconds int
	switch freq {
	case minute:
		freqInSeconds = 60
	case hour:
		freqInSeconds = 3600
	case day:
		freqInSeconds = 24 * 3600
	}

	var res []Interval
	startI := start
	for startI < end {
		nextFreqStart := ((startI + freqInSeconds) / freqInSeconds) * freqInSeconds
		endI := int(math.Min(float64(end), float64(nextFreqStart)))
		res = append(res, Interval{startI, endI})
		startI = endI
	}

	return res
}
