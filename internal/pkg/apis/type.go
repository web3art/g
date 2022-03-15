package apis

import (
	"fmt"

	"github.com/web3art/g/internal/pkg/dao/model"
)

type Twteet struct {
	ID                    string `json:"id"`
	Text                  string `json:"text"`
	CreatedAt             int64  `json:"createdAt"`
	AuthorId              string `json:"authorId"`
	AuthorName            string `json:"authorName"`
	AuthorUsername        string `json:"authorUsername"`
	Likes                 int    `json:"likes"`
	Retweets              int    `json:"retweets"`
	Score                 int    `json:"score"`
	Link                  string `json:"link"`
	AuthorProfileImageURL string `json:"authorProfileImageURL"`
}

func modelTweetToAPITwteet(in model.Tweet) Twteet {
	return Twteet{
		ID:                    fmt.Sprintf("%d", in.Id),
		Text:                  in.Text,
		CreatedAt:             in.CreatedAt.Unix(),
		AuthorId:              fmt.Sprintf("%d", in.AuthorId),
		AuthorName:            in.AuthorName,
		AuthorUsername:        in.AuthorUserName,
		Likes:                 in.LikeCount,
		Retweets:              in.RetweetCount,
		Score:                 in.Score,
		Link:                  fmt.Sprintf("https://twitter.com/%s/status/%d", in.AuthorUserName, in.Id),
		AuthorProfileImageURL: in.AuthorproFileImageUrl,
	}
}

type TemporaryToken struct {
	TokenId uint   `json:"tokenId"`
	Claimed bool   `json:"claimed"`
	Twteet  Twteet `json:"twteet"`
}

type TwteetToWin struct {
	TokenId *uint  `json:"tokenId,omitempty"`
	Claimed *bool  `json:"claimed,omitempty"`
	Twteet  Twteet `json:"twteet"`
}
