package tw

import (
	"fmt"
	"net/http"
)

type Tweet struct {
	token string
}

type authorize struct {
	Token string
}

func (a authorize) Add(req *http.Request) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.Token))
}

func MakeTweet(token string) *Tweet {
	return &Tweet{token: token}
}
