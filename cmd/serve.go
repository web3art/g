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

	tweet := tw.MakeTweet("AAAAAAAAAAAAAAAAAAAAAEbNXwEAAAAAFI8uK2vtlD1RbBp4OkHRbv3qxC0%3DgL16PjF1qoUXxaeRLntPAp0yWjCWfNjh1qBZX6ze51sDyOAY5g")
	tweet.SyncRecentPool([]string{"web3sword"})
	tweet.UpdateTweetScorePool()

	mux := http.NewServeMux()
	mux.HandleFunc("/recent-twteet", apis.RecentTwteet)
	mux.HandleFunc("/twteet-to-wins", apis.GetTwteetToWinList)
	mux.HandleFunc("/temporary-token", apis.GetTemporaryToken)

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST). See
	// documentation below for more options.
	handler := cors.Default().Handler(mux)
	http.ListenAndServe(fmt.Sprintf(":%s", port), handler)
	return nil
}
