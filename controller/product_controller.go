package controller

import (
	"net/http"
	"strconv"

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
	err := request.ParseForm()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}

	price, err := strconv.Atoi(request.FormValue("price"))
	if err != nil {
		panic(err)
	}
	quantity, err := strconv.Atoi(request.FormValue("quantity"))
	if err != nil {
		panic(err)
	}

	products := &entity.Product{
		Name:     request.FormValue("name"),
		Price:    float32(price),
		Quantity: quantity,
	}
	err = pc.ProductUsecase.Create(request.Context(), products)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}
