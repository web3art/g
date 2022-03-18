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
	7000012,
	8000012,
	7000013,
	8000013,
	5000019,
	6000019,
	9000019,
	10000019,
	8000023,
	9000023,
	8000024,
	9000024,
	6000027,
	7000027,
	6000028,
	7000028,
	8000031,
	9000031,
	8000032,
	9000032,
	6000035,
	7000035,
	6000036,
	7000036,
	8000039,
	9000039,
	8000040,
	9000040,
	6000043,
	7000043,
	6000044,
	7000044,
	8000046,
	7000047,
	8000047,
}

var twteetLucyTokenIds []uint = []uint{
	5000020,
	6000020,
	7000020,
	8000020,
	9000020,
	10000020,
	8000021,
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

	for _, t := range twteetToWinTokenIds {
		if tokenAssigned[t] {
			continue
		}
		var twteet model.Tweet
		if err := db.DB().Model(model.Tweet{}).Where("is_lucky_tweet = false and is_claim_tweet = false and assigned = false and author_id not in (?) and id not in (?)", db.DB().Table("tweet_wait_to_claims").Select("author_id"), twteetIds).Order("score desc").Limit(1).First(&twteet).Error; err != nil {
			fmt.Println("fetch get tweet error:", err.Error())
			continue
		}
		twteetIds = append(twteetIds, twteet.Id)
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
		if err := db.DB().Model(model.Tweet{}).Where("is_lucky_tweet = true and is_claim_tweet = false and assigned = false and author_id not in (?) and id not in (?)", db.DB().Table("tweet_wait_to_claims").Select("author_id"), twteetIds).Order("id desc").Limit(1).First(&twteet).Error; err != nil {
			fmt.Println("fetch get tweet error:", err.Error())
			continue
		}
		twteetIds = append(twteetIds, twteet.Id)
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
