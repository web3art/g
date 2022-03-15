package apis

import (
	"net/http"

	"github.com/web3art/g/internal/pkg/thegraph/sword"
)

func GetCurrentPrice(w http.ResponseWriter, r *http.Request) {
	q, err := sword.GetCurrentPrice()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(q.CurrentPrice.Price))
}
