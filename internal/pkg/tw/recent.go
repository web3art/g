package tw

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	gt "github.com/g8rswimmer/go-twitter/v2"
	"github.com/web3art/g/internal/pkg/dao/model"
	"github.com/web3art/g/internal/pkg/db"
	"github.com/web3art/g/internal/pkg/types"
	"gorm.io/gorm"
)

// Find recent tweets that are regular topics and not retweets
func (t *Tweet) SyncRecentSearchByTopicAndNoRetweet(topics []string) error {
	client := &gt.Client{
		Authorizer: authorize{
			Token: t.token,
		},
		Client: http.DefaultClient,
		Host:   "https://api.twitter.com",
	}

	var maxTwtterId uint
	if err := db.DB().Model(&model.Tweet{}).Order("id desc").Limit(1).Select("id").Scan(&maxTwtterId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			maxTwtterId = 0
		} else {
			return err
		}
	}

	opts := gt.TweetRecentSearchOpts{
		Expansions: []gt.Expansion{
			gt.ExpansionEntitiesMentionsUserName,
			gt.ExpansionAuthorID,
		},
		TweetFields: []gt.TweetField{
			gt.TweetFieldCreatedAt,
			gt.TweetFieldConversationID,
			gt.TweetFieldPublicMetrics,
			gt.TweetFieldInReplyToUserID,
		},
		UserFields: []gt.UserField{
			gt.UserFieldProfileImageURL,
		},
		SinceID: fmt.Sprintf("%d", maxTwtterId),
	}
	topichash := ""
	for _, topic := range topics {
		topichash += fmt.Sprintf("#%s ", topic)
	}
	query := fmt.Sprintf("%s -is:retweet", topichash)
	tweetResponse, err := client.TweetRecentSearch(context.Background(), query, opts)
	if err != nil {
		return err
	}
	dictionaries := tweetResponse.Raw.TweetDictionaries()
	for _, twteet := range dictionaries {
		if err := db.DB().First(&model.Tweet{}, "id = ?", twteet.Tweet.ID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				id, err := types.ParseInt64(twteet.Tweet.ID)
				if err != nil {
					return err
				}
				authorId, err := types.ParseInt64(twteet.Tweet.AuthorID)
				if err != nil {
					return err
				}
				db.DB().Create(&model.Tweet{
					Id:                    uint(id),
					AuthorId:              uint(authorId),
					AuthorName:            twteet.Author.Name,
					AuthorUserName:        twteet.Author.UserName,
					AuthorproFileImageUrl: twteet.Author.ProfileImageURL,
					LikeCount:             twteet.Tweet.PublicMetrics.Likes,
					RetweetCount:          twteet.Tweet.PublicMetrics.Retweets,
					Assigned:              false,
					// Retweets * 40% + Likes * 60%
					Score:        int(twteet.Tweet.PublicMetrics.Likes*40/100 + twteet.Tweet.PublicMetrics.Retweets*60/100),
					Text:         twteet.Tweet.Text,
					IsClaimTweet: strings.Contains(twteet.Tweet.Text, "#claim"),
					CreatedAt:    time.Now().UTC(),
					UpdatedAt:    time.Now().UTC(),
				})
			}
		}
	}
	return nil
}
