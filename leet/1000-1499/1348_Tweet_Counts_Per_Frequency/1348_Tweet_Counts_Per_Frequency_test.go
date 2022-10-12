package _1348_Tweet_Counts_Per_Frequency

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	tw := Constructor()
	tw.RecordTweet("tweet3", 0)
	tw.RecordTweet("tweet3", 60)
	tw.RecordTweet("tweet3", 10)

	res := tw.GetTweetCountsPerFrequency("minute", "tweet3", 0, 59)
	t.Log(res)

	res = tw.GetTweetCountsPerFrequency("minute", "tweet3", 0, 60)
	t.Log(res)

	tw.RecordTweet("tweet3", 120)

	res = tw.GetTweetCountsPerFrequency("hour", "tweet3", 0, 210)
	t.Log(res)
}
