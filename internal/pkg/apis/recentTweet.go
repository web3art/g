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

	if err := db.DB().Model(&model.Tweet{}).Order("created_at desc").Limit(int(first)).Where("is_claim_tweet = ?", false).Find(&tweets).Error; err != nil {
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

	if err := db.DB().Model(&model.Tweet{}).Order("score desc").Where("is_claim_tweet = ?", false).Find(&tweets).Error; err != nil {
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
	70012,
	80012,
	70013,
	80013,
	50019,
	60019,
	90019,
	100019,
	80023,
	90023,
	80024,
	90024,
	60027,
	70027,
	60028,
	70028,
	80031,
	90031,
	80032,
	90032,
	60035,
	70035,
	60036,
	70036,
	80039,
	90039,
	80040,
	90040,
	60043,
	70043,
	60044,
	70044,
	80046,
	70047,
	80047,
}

var twteetLucyTokenIds []uint = []uint{
	50020,
	60020,
	70020,
	80020,
	90020,
	100020,
	80021,
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
	authorIds := []uint{}

	for _, t := range twteetToWinTokenIds {
		if tokenAssigned[t] {
			continue
		}
		var twteet model.Tweet
		if err := db.DB().Model(model.Tweet{}).Where("is_lucky_tweet = false and is_claim_tweet = false and assigned = false and author_id not in (?) and author_id not in (?) and id not in (?)", db.DB().Table("tweet_wait_to_claims").Select("author_id"), authorIds, twteetIds).Order("score desc").Limit(1).First(&twteet).Error; err != nil {
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
		if err := db.DB().Model(model.Tweet{}).Where("is_lucky_tweet = true and is_claim_tweet = false and assigned = false and author_id not in (?) and author_id not in (?) and id not in (?)", db.DB().Table("tweet_wait_to_claims").Select("author_id"), authorIds, twteetIds).Order("id desc").Limit(1).First(&twteet).Error; err != nil {
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
