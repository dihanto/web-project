package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/dihanto/crud-web/entity"
	"github.com/dihanto/crud-web/usecase"
	"github.com/julienschmidt/httprouter"
)

type ProductController struct {
	ProductUsecase usecase.ProductUsecase
}

func NewProductController(productUsecase usecase.ProductUsecase) *ProductController {
	return &ProductController{
		ProductUsecase: productUsecase,
	}
}

func (pc *ProductController) Create(writer http.ResponseWriter, request *http.Request, p httprouter.Params) {
	err := request.ParseForm()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
	price, err := strconv.Atoi(request.FormValue("price"))
	if err != nil {
		panic(err)
	}
	log.Println(price)
	quantity, err := strconv.Atoi(request.FormValue("quantity"))
	if err != nil {
		panic(err)
	}
	name := request.FormValue("name")

	products := &entity.Product{
		Name:     name,
		Price:    float32(price),
		Quantity: quantity,
	}
	err = pc.ProductUsecase.Create(request.Context(), products)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	successCreate(writer, request)
	writer.WriteHeader(http.StatusCreated)
}

func (pc *ProductController) GetAll(writer http.ResponseWriter, request *http.Request, p httprouter.Params) {
	result, err := pc.ProductUsecase.GetAll(request.Context())
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
	writer.WriteHeader(http.StatusOK)

	t := template.Must(template.ParseFiles("./views/home/index.html"))
	if err != nil {
		log.Fatal(err)
	}
	if err := t.Execute(writer, result); err != nil {
		log.Fatal(err)
	}

}
func (pc *ProductController) FindById(writer http.ResponseWriter, request *http.Request, p httprouter.Params) {

	idstring := p.ByName("id")

	log.Println(idstring)

	id, err := strconv.Atoi(idstring)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}

	result, err := pc.ProductUsecase.FindById(request.Context(), id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	writer.WriteHeader(http.StatusOK)

	t := template.Must(template.ParseFiles("./views/product/findbyidresult.html"))
	if err != nil {
		log.Fatal(err)
	}
	if err := t.Execute(writer, result); err != nil {
		log.Fatal(err)
	}
}

func (pc *ProductController) Update(writer http.ResponseWriter, request *http.Request, p httprouter.Params) {
	idstring := p.ByName("id")

	log.Println(idstring)

	id, err := strconv.Atoi(idstring)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}

	var data struct {
		Name     string `json:"name"`
		Price    string `json:"price"`
		Quantity string `json:"quantity"`
	}

	err = json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}

	price, err := strconv.Atoi(data.Price)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
	quantity, err := strconv.Atoi(data.Quantity)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}

	products := &entity.Product{
		ID:       int64(id),
		Name:     data.Name,
		Price:    float32(price),
		Quantity: quantity,
	}
	log.Println(products)
	err = pc.ProductUsecase.Update(request.Context(), products)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	successUpdate(writer, request)
	writer.WriteHeader(http.StatusOK)
}

func (pc *ProductController) Delete(writer http.ResponseWriter, request *http.Request, p httprouter.Params) {

	idstring := p.ByName("id")
	log.Println(idstring)
	id, err := strconv.Atoi(idstring)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}

	err = pc.ProductUsecase.Delete(request.Context(), id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	successDelete(writer, request)
}

func successCreate(writer http.ResponseWriter, request *http.Request) {

	tmpl, err := template.ParseFiles("views/product/successcreate.html")
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

func successUpdate(writer http.ResponseWriter, request *http.Request) {

	tmpl, err := template.ParseFiles("views/product/successupdate.html")
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

func successDelete(writer http.ResponseWriter, request *http.Request) {

	tmpl, err := template.ParseFiles("views/product/successdelete.html")
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

func ProductIndex(writer http.ResponseWriter, request *http.Request) {

	tmpl, err := template.ParseFiles("views/home/index.html")
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

func (pc *ProductController) ProductEdit(writer http.ResponseWriter, request *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := pc.ProductUsecase.FindById(request.Context(), id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
	tmpl, err := template.ParseFiles("views/product/update.html")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(writer, result[0])
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
func (pc *ProductController) ProductAdd(writer http.ResponseWriter, request *http.Request, p httprouter.Params) {

	tmpl, err := template.ParseFiles("views/product/add.html")
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
