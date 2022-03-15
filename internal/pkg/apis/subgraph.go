package apis

import (
	"encoding/json"
	"net/http"
	"os"
)

func GetConfig(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	subgraphEndpoint := os.Getenv("SUBGRAPH_ENDPOINT")
	swordAddress := os.Getenv("SWORD_ADDRESS")
	ethNetwork := os.Getenv("ETH_NETWORK")

	config := Config{
		SubgraphEndpoint: subgraphEndpoint,
		SowrdAddress:     swordAddress,
		EthNetwork:       ethNetwork,
	}

	j, err := json.Marshal(config)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(j)
}
