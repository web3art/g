package sword

import (
	"context"
	"os"

	"github.com/hasura/go-graphql-client"
)

// {
// 	currentUnlockTokenId(id: 6) {
// 	  tokenIds
// 	}
// }
type UnlockTwteetTokenQuery struct {
	CurrentUnlockTokenId struct {
		TokenIds []string `json:"tokenIds"`
	} `graphql:"currentUnlockTokenId(id: 6)"`
}

func GetUnlockTwteetTokenIds() ([]string, error) {
	endpint := os.Getenv("SUBGRAPH_ENDPOINT")
	client := graphql.NewClient(endpint, nil)

	var q UnlockTwteetTokenQuery

	if err := client.Query(context.Background(), &q, nil); err != nil {
		return nil, err
	}

	return q.CurrentUnlockTokenId.TokenIds, nil
}
