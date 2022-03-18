package tw

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"

	gt "github.com/g8rswimmer/go-twitter/v2"
	"github.com/web3art/g/internal/pkg/dao/model"
	"github.com/web3art/g/internal/pkg/db"
	"github.com/web3art/g/internal/pkg/types"
	"gorm.io/gorm"
)

// I'm eth address is 0x714df076992f95E452A345cD8289882CEc6ab82F
func ExtractEthAddress(text string) (*string, error) {
	index := strings.Index(text, "0x")

	if index == -1 {
		return nil, fmt.Errorf("no eth address found")
	}

	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	address := text[index : index+42]

	// check if address is valid
	if !re.MatchString(address) {
		return nil, fmt.Errorf("invalid eth address")
	}

	return &address, nil
}

// Find recent tweets that are regular topics and not retweets
func (t *Tweet) SyncRecentSearchByTopicAndNoRetweet(topics []string) error {
	client := &gt.Client{
		Authorizer: authorize{
			Token: t.token,
		},
		Client: http.DefaultClient,
		Host:   "https://api.twitter.com",
	}

	var maxTwtterId uint = 1503444294723645447
	if err := db.DB().Model(&model.Tweet{}).Order("id desc").Limit(1).Select("id").Scan(&maxTwtterId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			maxTwtterId = 1503444294723645447
		} else {
			return err
		}
	}

	opts := gt.TweetRecentSearchOpts{
		Expansions: []gt.Expansion{
			gt.ExpansionEntitiesMentionsUserName,
			gt.ExpansionAuthorID,
			gt.ExpansionAttachmentsMediaKeys,
		},
		TweetFields: []gt.TweetField{
			gt.TweetFieldCreatedAt,
			gt.TweetFieldConversationID,
			gt.TweetFieldPublicMetrics,
		},
		UserFields: []gt.UserField{
			gt.UserFieldProfileImageURL,
		},
		MediaFields: []gt.MediaField{
			gt.MediaFieldURL,
			gt.MediaFieldPreviewImageURL,
		},
		SinceID: fmt.Sprintf("%d", maxTwtterId),
	}

	topichash := ""
	for _, topic := range topics {
		topichash += fmt.Sprintf("#%s ", topic)
	}
	query := fmt.Sprintf(strings.TrimRight("%s -is:retweet", " "), topichash)

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

				IsLuckyTweet := strings.Contains(twteet.Tweet.Text, "#LuckyW3S")
				isClaimTweet := strings.Contains(twteet.Tweet.Text, "#claim")
				TwteetImageURL := ""

				if isClaimTweet {
					// parse tweet eth address insert to Model.TweetAuthorAddress
					if ethAddress, err := ExtractEthAddress(twteet.Tweet.Text); err != nil {
						return err
					} else {
						if err := db.DB().Create(&model.TweetAuthorAddress{
							BindTweetId: uint(id),
							AuthorId:    uint(authorId),
							Address:     *ethAddress,
						}).Error; err != nil {
							return err
						}
					}
				}

				for _, media := range twteet.AttachmentMedia {
					if media.Type == "photo" {
						TwteetImageURL = media.URL
					}
				}

				db.DB().Create(&model.Tweet{
					Id:                    uint(id),
					AuthorId:              uint(authorId),
					AuthorName:            twteet.Author.Name,
					AuthorUserName:        twteet.Author.UserName,
					AuthorproFileImageUrl: twteet.Author.ProfileImageURL,
					TwteetImageURL:        TwteetImageURL,
					LikeCount:             twteet.Tweet.PublicMetrics.Likes,
					RetweetCount:          twteet.Tweet.PublicMetrics.Retweets,
					Assigned:              false,
					// Retweets * 40% + Likes * 60%
					Score:        int(twteet.Tweet.PublicMetrics.Likes*40/100 + twteet.Tweet.PublicMetrics.Retweets*60/100),
					Text:         twteet.Tweet.Text,
					IsClaimTweet: isClaimTweet,
					IsLuckyTweet: IsLuckyTweet,
					CreatedAt:    time.Now().UTC(),
					UpdatedAt:    time.Now().UTC(),
				})
			}
		}
	}
	return nil
}
