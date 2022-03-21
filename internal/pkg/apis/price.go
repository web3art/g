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

	if q.CurrentPrice.Price == "" {
		q.CurrentPrice.Price = "150000000000000000"
	}

	w.Write([]byte(q.CurrentPrice.Price))
}
