package model

import (
	"time"
)

type Tweet struct {
	Id                    uint   `gorm:"primaryKey"`
	AuthorId              uint   `gorm:"index"`
	AuthorName            string `gorm:"type:varchar(170)"`
	AuthorUserName        string `gorm:"type:varchar(170)"`
	AuthorproFileImageUrl string `gorm:"type:varchar(255)"`
	TwteetImageURL        string `gorm:"type:varchar(255)"`
	LikeCount             int
	RetweetCount          int
	Score                 int `gorm:"index"`
	Assigned              bool
	Text                  string
	IsClaimTweet          bool
	IsLuckyTweet          bool
	IsAirdropTweet        bool
	CreatedAt             time.Time `gorm:"index"`
	UpdatedAt             time.Time `gorm:"index"`
}

type TweetAuthorAddress struct {
	Id          uint `gorm:"primaryKey;AUTO_INCREMENT"`
	AuthorId    uint `gorm:"index"`
	Address     string
	BindTweetId uint
}

type TweetWaitToClaim struct {
	Id       uint `gorm:"primaryKey;AUTO_INCREMENT"`
	TweetId  uint `gorm:"uniqueIndex"`
	AuthorId uint `gorm:"uniqueIndex"`
	TokenId  uint `gorm:"uniqueIndex"`
	Claimed  bool
}

type UserStatistics struct {
	Id            uint `gorm:"primaryKey;AUTO_INCREMENT"`
	UserJoinCount uint
}

type BlacklistedUsers struct {
	AuthorId uint `gorm:"primaryKey;"`
}
