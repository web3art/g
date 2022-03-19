package tw

import (
	"context"
	"fmt"
	"net/http"
	"time"

	gt "github.com/g8rswimmer/go-twitter/v2"
	"github.com/web3art/g/internal/pkg/dao/model"
	"github.com/web3art/g/internal/pkg/db"
)

func (t *Tweet) UpdateTweetScore() error {
	var tweets []model.Tweet
	if err := db.DB().Model(&model.Tweet{}).Order("updated_at asc").Limit(200).Where("is_claim_tweet = ?", false).Find(&tweets).Error; err != nil {
		return err
	}

	client := &gt.Client{
		Authorizer: authorize{
			Token: t.token,
		},
		Client: http.DefaultClient,
		Host:   "https://api.twitter.com",
	}

	opts := gt.TweetLookupOpts{
		Expansions: []gt.Expansion{gt.ExpansionEntitiesMentionsUserName, gt.ExpansionAuthorID},
		TweetFields: []gt.TweetField{
			gt.TweetFieldPublicMetrics,
		},
	}
	for _, tweet := range tweets {
		tweetResponse, err := client.TweetLookup(context.Background(), []string{
			fmt.Sprintf("%d", tweet.Id),
		}, opts)
		if err != nil {
			return err
		}
		if tweetResponse.Raw.Errors != nil {
			return fmt.Errorf("%s", tweetResponse.Raw.Errors[0].Detail)
		}

		dictionaries := tweetResponse.Raw.TweetDictionaries()
		for _, t := range dictionaries {
			tweet.LikeCount = t.Tweet.PublicMetrics.Likes
			tweet.RetweetCount = t.Tweet.PublicMetrics.Retweets
			tweet.UpdatedAt = time.Now().UTC()
			tweet.Score = int(t.Tweet.PublicMetrics.Likes*40/100 + t.Tweet.PublicMetrics.Retweets*60/100)
			if err := db.DB().Save(&tweet).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
