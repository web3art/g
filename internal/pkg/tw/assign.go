package tw

import (
	"fmt"
	"time"

	"github.com/web3art/g/internal/pkg/dao/model"
	"github.com/web3art/g/internal/pkg/db"
	"github.com/web3art/g/internal/pkg/thegraph/sword"
	"github.com/web3art/g/internal/pkg/types"
	"gorm.io/gorm"
)

// Periodically get the unlocked twteet token id in order to give it to the person with the highest current score
func AssignTwteetToWinWaitToClaimTokenPool() {
	go func() {
		for {
			time.Sleep(5 * time.Minute)

			tokenIds, err := sword.GetUnlockTwteetTokenIds()

			if err != nil {
				fmt.Println("fetch token ids error:", err.Error())
			}

			for _, tokenId := range tokenIds {
				intTokenId, err := types.ParseInt64(tokenId)
				if err != nil {
					fmt.Println("parse token id error:", err.Error())
					continue
				}
				uintTokenId := uint(intTokenId)
				err = db.DB().Model(model.TweetWaitToClaim{}).Where("token_id = ?", tokenId).First(&model.TweetWaitToClaim{}).Error
				if err != nil && err != gorm.ErrRecordNotFound {
					fmt.Println("get token error:", err.Error())
					continue
				} else if err == gorm.ErrRecordNotFound {
					// fetch twteet score highest
					var twteet model.Tweet
					if err := db.DB().Model(model.Tweet{}).Where("is_airdrop_tweet = false and is_lucky_tweet = false and is_claim_tweet = false and assigned = false and score > 0 and author_id not in (?)", db.DB().Table("tweet_wait_to_claims").Select("author_id")).Order("score desc, retweet_count desc, like_count desc").Limit(1).First(&twteet).Error; err != nil {
						fmt.Println("fetch get tweet error:", err.Error())
						continue
					}
					db.DB().Transaction(func(tx *gorm.DB) error {
						twteet.Assigned = true
						if err := tx.Save(&twteet).Error; err != nil {
							return err
						}
						twtc := model.TweetWaitToClaim{
							TweetId:  twteet.Id,
							Claimed:  false,
							TokenId:  uintTokenId,
							AuthorId: twteet.AuthorId,
						}
						if err := tx.Create(&twtc).Error; err != nil {
							return err
						}
						return nil
					})
				} else {
					// fmt.Println("token already exists:", tokenId)
					continue
				}
			}
		}
	}()
}

func AssignLuckyWaitToClaimTokenPool() {
	go func() {
		for {
			time.Sleep(5 * time.Minute)

			tokenIds, err := sword.GetUnlockLuckyTokenIds()

			if err != nil {
				fmt.Println("fetch token ids error:", err.Error())
			}

			for _, tokenId := range tokenIds {
				intTokenId, err := types.ParseInt64(tokenId)
				if err != nil {
					fmt.Println("parse token id error:", err.Error())
					continue
				}
				uintTokenId := uint(intTokenId)
				err = db.DB().Model(model.TweetWaitToClaim{}).Where("token_id = ?", tokenId).First(&model.TweetWaitToClaim{}).Error
				if err != nil && err != gorm.ErrRecordNotFound {
					fmt.Println("get token error:", err.Error())
					continue
				} else if err == gorm.ErrRecordNotFound {
					// fetch twteet score highest
					var twteet model.Tweet
					if err := db.DB().Model(model.Tweet{}).Where("is_lucky_tweet = true and is_claim_tweet = false and assigned = false and score > 0 and author_id not in (?)", db.DB().Table("tweet_wait_to_claims").Select("author_id")).Order("RAND()").Limit(1).First(&twteet).Error; err != nil {
						fmt.Println("fetch get tweet error:", err.Error())
						continue
					}
					db.DB().Transaction(func(tx *gorm.DB) error {
						twteet.Assigned = true
						if err := tx.Save(&twteet).Error; err != nil {
							return err
						}
						twtc := model.TweetWaitToClaim{
							TweetId:  twteet.Id,
							Claimed:  false,
							TokenId:  uintTokenId,
							AuthorId: twteet.AuthorId,
						}
						if err := tx.Create(&twtc).Error; err != nil {
							return err
						}
						return nil
					})
				} else {
					// fmt.Println("token already exists:", tokenId)
					continue
				}
			}
		}
	}()
}
