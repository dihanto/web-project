package router

import (
	"net/http"
	"os"
	"text/template"

	"github.com/dihanto/crud-web/controller"
	"github.com/dihanto/crud-web/middleware"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(pc *controller.ProductController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/products", func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		create(writer, request)
	})

	router.POST("/products", middleware.CheckHandler(func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		pc.Create(w, r)
	}))

	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		home(w, r)
	})

	return router
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

func create(writer http.ResponseWriter, request *http.Request) {

	tmpl, err := template.ParseFiles("views/product/index.html")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(writer, nil)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
