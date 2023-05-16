package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

func CheckHandler(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		name := r.FormValue("name")
		price, err := strconv.Atoi(r.FormValue("price"))
		if err != nil {
			panic(err)
		}
		quantity, err := strconv.Atoi(r.FormValue("quantity"))
		if err != nil {
			panic(err)
		}

		switch {
		case name == "":
			fmt.Fprintf(w, "name should not be empty")
		case price == 0 || price < 0:
			fmt.Fprintf(w, "price should be greater than 0")
		case quantity == 0 || quantity < 0:
			fmt.Fprintf(w, "price should be greater than 0")
		}

		h(w, r, p)
	}
}

func Logger(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		start := time.Now()

		h(w, r, p)

		log.Printf(
			"%s %s %s",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
	}
}
