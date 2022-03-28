package apis

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/web3art/g/internal/pkg/dao/model"
	"github.com/web3art/g/internal/pkg/db"
	"github.com/web3art/g/internal/pkg/types"
	"gorm.io/gorm"
)

func RecentTwteet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	first, err := types.ParseInt64(r.URL.Query().Get("first"))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"error\": \"invalid first\"}"))
		return
	}

	var tweets []model.Tweet

	if err := db.DB().Model(&model.Tweet{}).Order("created_at desc").Limit(int(first)).Where("is_claim_tweet = ? and author_id not in (?)", false, db.DB().Table("blacklisted_users").Select("author_id")).Find(&tweets).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	var mts []Twteet
	for _, t := range tweets {
		mts = append(mts, modelTweetToAPITwteet(t))
	}

	j, err := json.Marshal(mts)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(j)
}

func GetTwteetToWinList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var tweets []model.Tweet

	if err := db.DB().Model(&model.Tweet{}).Order("score desc, retweet_count desc, like_count desc").Where("is_claim_tweet = ?", false).Find(&tweets).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	var ttws []TwteetToWin = []TwteetToWin{}

	for _, t := range tweets {
		var twtc *model.TweetWaitToClaim
		if err := db.DB().Where("tweet_id = ?", t.Id).First(&twtc).Error; err != nil && err != gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		if twtc != nil {
			ttws = append(ttws, TwteetToWin{
				Twteet:  modelTweetToAPITwteet(t),
				TokenId: &twtc.TokenId,
				Claimed: &twtc.Claimed,
			})
		} else {
			ttws = append(ttws, TwteetToWin{
				Twteet: modelTweetToAPITwteet(t),
			})
		}
	}

	j, err := json.Marshal(ttws)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(j)
}

var twteetToWinTokenIds []uint = []uint{
	80013,
	90013,
	80014,
	90014,
	60020,
	70020,
	100020,
	110020,
	90024,
	100024,
	90025,
	100025,
	70028,
	80028,
	70029,
	80029,
	90032,
	100032,
	90033,
	100033,
	70036,
	80036,
	70037,
	80037,
	90040,
	100040,
	90041,
	100041,
	70044,
	80044,
	70045,
	80045,
	90047,
	80048,
	90048,
}

var twteetLucyTokenIds []uint = []uint{
	60021,
	70021,
	80021,
	90021,
	100021,
	110021,
	90022,
}

var airdropTokenIds []uint = []uint{
	90028,
	100028,
	110028,
	90029,
	100029,
	110029,
	90030,
	100030,
	110030,
	90031,
	100031,
	110031,
	60032,
	70032,
	80032,
	70033,
	80033,
	70034,
	80034,
	70035,
	80035,
}

func GetTemporaryToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var ts []TemporaryToken = []TemporaryToken{}
	var tcs []model.TweetWaitToClaim

	if err := db.DB().Model(&model.TweetWaitToClaim{}).Order("id desc").Limit(100).Find(&tcs).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	tokenAssigned := map[uint]bool{}

	for _, t := range tcs {
		// TokenId int64  `json:"tokenId"`
		// Claimed bool   `json:"claimed"`
		// Twteet  Twteet `json:"twteet"`
		var tw model.Tweet
		if err := db.DB().First(&tw, t.TweetId).Error; err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		ts = append(ts, TemporaryToken{
			TokenId:    t.TokenId,
			Claimed:    t.Claimed,
			CanClaimed: true,
			Twteet:     modelTweetToAPITwteet(tw),
		})
		tokenAssigned[t.TokenId] = true
	}

	twteetIds := []uint{
		0,
	}
	authorIds := []uint{
		0,
	}

	for _, t := range twteetToWinTokenIds {
		if tokenAssigned[t] {
			continue
		}
		var twteet model.Tweet
		if err := db.DB().Model(model.Tweet{}).Where("is_airdrop_tweet = false and is_lucky_tweet = false and is_claim_tweet = false and assigned = false and author_id not in (?) and author_id not in (?) and id not in (?) and author_id not in (?)", db.DB().Table("tweet_wait_to_claims").Select("author_id"), authorIds, twteetIds, db.DB().Table("blacklisted_users").Select("author_id")).Order("score desc, retweet_count desc, like_count desc").Limit(1).First(&twteet).Error; err != nil {
			fmt.Println("fetch get tweet error:", err.Error())
			continue
		}
		twteetIds = append(twteetIds, twteet.Id)
		authorIds = append(authorIds, twteet.AuthorId)
		ts = append(ts, TemporaryToken{
			TokenId:    t,
			Claimed:    false,
			CanClaimed: false,
			Twteet:     modelTweetToAPITwteet(twteet),
		})
	}

	for _, t := range twteetLucyTokenIds {
		if tokenAssigned[t] {
			continue
		}
		var twteet model.Tweet
		if err := db.DB().Model(model.Tweet{}).Where("is_lucky_tweet = true and is_claim_tweet = false and assigned = false and author_id not in (?) and author_id not in (?) and id not in (?) and author_id not in (?)", db.DB().Table("tweet_wait_to_claims").Select("author_id"), authorIds, twteetIds, db.DB().Table("blacklisted_users").Select("author_id")).Order("id desc").Limit(1).First(&twteet).Error; err != nil {
			fmt.Println("fetch get tweet error:", err.Error())
			continue
		}
		twteetIds = append(twteetIds, twteet.Id)
		authorIds = append(authorIds, twteet.AuthorId)
		ts = append(ts, TemporaryToken{
			TokenId:    t,
			Claimed:    false,
			CanClaimed: false,
			Twteet:     modelTweetToAPITwteet(twteet),
		})
	}

	for _, t := range airdropTokenIds {
		if tokenAssigned[t] {
			continue
		}
		var twteet model.Tweet
		if err := db.DB().Model(model.Tweet{}).Where("is_airdrop_tweet = true and is_claim_tweet = false and assigned = false and author_id not in (?) and author_id not in (?) and id not in (?) and author_id not in (?)", db.DB().Table("tweet_wait_to_claims").Select("author_id"), authorIds, twteetIds, db.DB().Table("blacklisted_users").Select("author_id")).Order("id desc").Limit(1).First(&twteet).Error; err != nil {
			fmt.Println("fetch get tweet error:", err.Error())
			continue
		}
		twteetIds = append(twteetIds, twteet.Id)
		authorIds = append(authorIds, twteet.AuthorId)
		ts = append(ts, TemporaryToken{
			TokenId:    t,
			Claimed:    false,
			CanClaimed: false,
			Twteet:     modelTweetToAPITwteet(twteet),
		})
	}

	j, err := json.Marshal(ts)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(j)
}
