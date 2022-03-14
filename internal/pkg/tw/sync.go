package tw

import (
	"fmt"
	"time"
)

// Pull the most recent tweets every 5 minutes
func (tw *Tweet) SyncRecentPool(topics []string) error {
	go func() {
		for {
			err := tw.SyncRecentSearchByTopicAndNoRetweet(topics)
			if err != nil {
				fmt.Printf("sync recent tweet error %s", err.Error())
			}
			time.Sleep(5 * time.Minute)
		}
	}()
	return nil
}

// Last 300 tweets score updated every 5 minutes
func (tw *Tweet) UpdateTweetScorePool() error {
	go func() {
		for {
			println("start update tweet score")
			err := tw.UpdateTweetScore()
			if err != nil {
				fmt.Printf("update tweet score error %s", err.Error())
			}
			time.Sleep(1 * time.Minute)
		}
	}()
	return nil
}
