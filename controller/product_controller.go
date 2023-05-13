package controller

import (
	"encoding/json"
	"net/http"
	"text/template"

	"github.com/dihanto/crud-web/entity"
	"github.com/dihanto/crud-web/usecase"
)

type ProductController struct {
	ProductUsecase usecase.ProductUsecase
}

func NewProductController(productUsecase usecase.ProductUsecase) *ProductController {
	return &ProductController{
		ProductUsecase: productUsecase,
	}
}

func (pc *ProductController) Create(writer http.ResponseWriter, request *http.Request) {

	if request.Method == "GET" {
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
	} else if request.Method == "POST" {
		products := &entity.Product{}
		err := json.NewDecoder(request.Body).Decode(products)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		err = pc.ProductUsecase.Create(request.Context(), products)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		writer.WriteHeader(http.StatusCreated)
	}
}
