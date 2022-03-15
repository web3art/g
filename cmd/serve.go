package cmd

import (
	"fmt"
	"os"

	"net/http"

	"github.com/rs/cors"
	"github.com/urfave/cli/v2"
	"github.com/web3art/g/internal/pkg/apis"
	"github.com/web3art/g/internal/pkg/tw"
)

func Serve(c *cli.Context) error {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	err := xinit(c)

	if err != nil {
		return err
	}

	tweet := tw.MakeTweet(os.Getenv("TWITTER_TOKEN"))
	tweet.SyncRecentPool([]string{"web3sword"})
	tweet.UpdateTweetScorePool()
	tw.AssignWaitToClaimTokenPool()

	mux := http.NewServeMux()
	mux.HandleFunc("/recent-twteet", apis.RecentTwteet)
	mux.HandleFunc("/twteet-to-wins", apis.GetTwteetToWinList)
	mux.HandleFunc("/temporary-token", apis.GetTemporaryToken)
	mux.HandleFunc("/current-price", apis.GetCurrentPrice)
	mux.HandleFunc("/config", apis.GetConfig)

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST). See
	// documentation below for more options.
	handler := cors.Default().Handler(mux)
	http.ListenAndServe(fmt.Sprintf(":%s", port), handler)
	return nil
}
