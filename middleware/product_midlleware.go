package middleware

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func CheckHandler(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		price, err := strconv.Atoi(r.FormValue("price"))
		if err != nil {
			panic(err)
		}

		if price == 0 || price < 0 {
			fmt.Fprintf(w, "price should be greater than 0")
			return
		}

		h(w, r, p)
	}
}
