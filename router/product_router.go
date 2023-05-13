package router

import (
	"net/http"
	"os"

	"github.com/dihanto/crud-web/controller"
	"github.com/gorilla/mux"
)

func NewRouter(pc *controller.ProductController) *mux.Router {
	r := mux.NewRouter()
	// r.Use(middleware.Authenticate)

	r.HandleFunc("/", home)
	r.HandleFunc("/products", pc.Create).Methods(http.MethodGet, http.MethodPost)
	return r
}

func home(writer http.ResponseWriter, request *http.Request) {
	file, err := os.ReadFile("views/home/index.html")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "text.html; charset=utf-8")
	writer.Write(file)
}
