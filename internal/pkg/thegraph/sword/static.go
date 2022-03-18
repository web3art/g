package sword

import (
	"context"
	"os"

	"github.com/hasura/go-graphql-client"
)

// type Static @entity {
// 	id: ID!
// 	revisionsoFar: BigInt!
// 	plotLeft: BigInt!
// }
type StaticQuery struct {
	Static struct {
		ID            string `json:"id"`
		RevisionsoFar string `json:"revisionsoFar"`
		PlotLeft      string `json:"plotLeft"`
	} `graphql:"static(id: 0)"`
}

func GetStatic() (*StaticQuery, error) {
	endpint := os.Getenv("SUBGRAPH_ENDPOINT")
	client := graphql.NewClient(endpint, nil)

	var q StaticQuery

	if err := client.Query(context.Background(), &q, nil); err != nil {
		return nil, err
	}

	return &q, nil
}
