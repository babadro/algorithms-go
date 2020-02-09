package _1348_Tweet_Counts_Per_Frequency

import (
	"testing"
)

func TestTweetCounts(t *testing.T) {
	tweetCounts := Constructor()
	tweetCounts.RecordTweet("tweet3", 0)
	tweetCounts.RecordTweet("tweet3", 59)
	tweetCounts.RecordTweet("tweet3", 60)
	tweetCounts.RecordTweet("tweet3", 10)
	tweetCounts.RecordTweet("tweet3", 120)

	//t.Log(tweetCounts.GetTweetCountsPerFrequency("minute", "tweet3", 0, 59))
	t.Log(tweetCounts.GetTweetCountsPerFrequency("minute", "tweet3", 0, 60))
	t.Log(tweetCounts.GetTweetCountsPerFrequency("hour", "tweet3", 0, 210))
	t.Log(tweetCounts.GetTweetCountsPerFrequency("minute", "tweet3", 59, 60))
}

func TestSmth(t *testing.T) {
	//arr := []int{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
	//t.Log(len(arr))
	//t.Log(len([]int{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}))
	twCounts := Constructor()
	twCounts.RecordTweet("tweet0", 19)
	twCounts.RecordTweet("tweet1", 34)
	twCounts.RecordTweet("tweet2", 36)
	twCounts.RecordTweet("tweet3", 59)
	twCounts.RecordTweet("tweet4", 64)
	twCounts.RecordTweet("tweet2", 48)
	twCounts.RecordTweet("tweet4", 21)
	twCounts.RecordTweet("tweet2", 44)

	t.Log(len(twCounts.GetTweetCountsPerFrequency("minute", "tweet1", 59, 9302)))

	twCounts.RecordTweet("tweet2", 74)

	t.Log(twCounts.GetTweetCountsPerFrequency("minute", "tweet4", 64, 2783))

	twCounts.RecordTweet("tweet2", 16)

}

/*
input:
["TweetCounts","recordTweet","recordTweet","recordTweet","recordTweet","recordTweet","recordTweet","recordTweet","recordTweet","getTweetCountsPerFrequency","recordTweet","getTweetCountsPerFrequency","recordTweet"]
[[],["tweet0",19],["tweet1",34],["tweet2",36],["tweet3",59],["tweet4",64],["tweet2",48],["tweet4",21],["tweet2",44],["minute","tweet1",59,9302],["tweet2",74],["minute","tweet4",64,2783],["tweet2",16]]

output/expected:
[null,null,null,null,null,null,null,null,null,[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],null,[1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],null]
[null,null,null,null,null,null,null,null,null,[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],null,[1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],null]
*/
