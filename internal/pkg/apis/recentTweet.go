package apis

import (
	"encoding/json"
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

	if err := db.DB().Model(&model.Tweet{}).Order("created_at desc").Where("is_claim_tweet = ?", false).Find(&tweets).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	var ttws []TwteetToWin = []TwteetToWin{}

	for _, t := range tweets {
		var twtc *model.TweetWaitToClaim
		if err := db.DB().First(&twtc, t.Id).Error; err != nil && err != gorm.ErrRecordNotFound {
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

func GetTemporaryToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var ts []TemporaryToken = []TemporaryToken{}
	var tcs []model.TweetWaitToClaim

	if err := db.DB().Model(&model.TweetWaitToClaim{}).Order("id desc").Limit(100).Find(&tcs).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

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
			TokenId: t.TokenId,
			Claimed: false,
			Twteet:  modelTweetToAPITwteet(tw),
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
