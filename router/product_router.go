package router

import (
	"net/http"
	"text/template"

	"github.com/dihanto/crud-web/controller"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(pc *controller.ProductController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/product", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		product(w, r)
	})
	router.POST("/product", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		name := r.FormValue("name")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		idupdate := r.FormValue("id")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		idDelete := r.FormValue("id_delete")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		switch {
		case r.Method == "POST" && name != "":
			pc.Create(w, r)
		case r.Method == "POST" && idupdate != "":
			pc.Update(w, r)
		case r.Method == "POST" && idDelete != "":
			pc.Delete(w, r)
		}

	})

	return router
}

func product(writer http.ResponseWriter, request *http.Request) {

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

//	func findbyid(writer http.ResponseWriter, request *http.Request) {
//		tmpl, err := template.ParseFiles("views/product/findbyid.html")
//		if err != nil {
//			http.Error(writer, err.Error(), http.StatusInternalServerError)
//			return
//		}
//		err = tmpl.Execute(writer, nil)
//		if err != nil {
//			http.Error(writer, err.Error(), http.StatusInternalServerError)
//			return
//		}
//	}
