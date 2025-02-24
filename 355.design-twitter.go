package main

import (
	"container/heap"
	"fmt"
)

type Twitter struct {
	followMap map[int]map[int]bool
	tweetMap  map[int][]Tweet
	now       int
}

type Tweet struct {
	tweetID   int
	timestamp int
}

type Feed []Tweet

func (p Feed) Len() int {
	return len(p)
}

func (p Feed) Less(i, j int) bool {
	return p[i].timestamp > p[j].timestamp
}

func (p Feed) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *Feed) Push(x any) {
	*p = append(*p, x.(Tweet))
}

func (p *Feed) Pop() any {
	old := *p
	l := old[len(old)-1]
	*p = old[:len(old)-1]
	return l
}

func Constructor() Twitter {
	return Twitter{
		followMap: make(map[int]map[int]bool),
		tweetMap:  make(map[int][]Tweet),
		now:       0,
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.now++
	t := Tweet{
		tweetID:   tweetId,
		timestamp: this.now,
	}
	this.tweetMap[userId] = append(this.tweetMap[userId], t)
	fmt.Println(this)
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	var followees []int
	for id, yes := range this.followMap[userId] {
		if yes {
			followees = append(followees, id)
		}
	}
	followees = append(followees, userId)

	var allTweets []Tweet
	for _, id := range followees {
		allTweets = append(allTweets, this.tweetMap[id]...)
	}

	feed := &Feed{}
	heap.Init(feed)
	for _, tweet := range allTweets {
		heap.Push(feed, tweet)
	}

	result := make([]int, 0, len(*feed))
	for range 10 {
		if len(*feed) == 0 {
			break
		}
		t := heap.Pop(feed).(Tweet)
		result = append(result, t.tweetID)
	}
	return result
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if followeeId == followerId {
		return
	}
	if this.followMap[followerId] == nil {
		this.followMap[followerId] = make(map[int]bool)
	}
	this.followMap[followerId][followeeId] = true
	fmt.Println(this)
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if this.followMap[followerId] == nil {
		return
	}
	this.followMap[followerId][followeeId] = false
	fmt.Println(this)
}
