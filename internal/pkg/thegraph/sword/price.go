package sword

import (
	"context"
	"os"

	"github.com/hasura/go-graphql-client"
)

// {
// 	currentPrice(id: 0) {
// 	  id
// 	  price
// 	}
// }
type CurrentPriceQuery struct {
	CurrentPrice struct {
		ID    string `json:"id"`
		Price string `json:"price"`
	} `graphql:"currentPrice(id: 0)"`
}

func GetCurrentPrice() (*CurrentPriceQuery, error) {
	endpint := os.Getenv("SUBGRAPH_ENDPOINT")
	client := graphql.NewClient(endpint, nil)

	var q CurrentPriceQuery

	if err := client.Query(context.Background(), &q, nil); err != nil {
		return nil, err
	}

	return &q, nil
}
