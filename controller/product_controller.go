package controller

import (
	"log"
	"net/http"
	"strconv"
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

func (pc *ProductController) GetAll(writer http.ResponseWriter, request *http.Request) {
	result, err := pc.ProductUsecase.GetAll(request.Context())
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
	writer.WriteHeader(http.StatusOK)

	tmpl := `<html>
	<head><title>Product</title></head>
	<body>
	  <table>
		<tr><th>ID</th><th>Name</th><th>Price</th></tr>
		{{range .}}
		  <tr> <td>{{.ID}}</td><td>{{.Name}}</td><td>{{.Price}}</td>
		  </tr>
		{{end}}
	  </table>
	</body>
  </html>`
	t, err := template.New("product").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	if err := t.Execute(writer, result); err != nil {
		log.Fatal(err)
	}

}
func (pc *ProductController) FindById(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
	id, err := strconv.Atoi(request.FormValue("id"))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
	result, err := pc.ProductUsecase.FindById(request.Context(), id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	writer.WriteHeader(http.StatusOK)

	tmpl := `<html>
	<head><title>Product</title></head>
	<body>
	  <table>
		<tr><th>ID</th><th>Name</th><th>Price</th><th>Quantity</tr>
		{{range .}}
		  <tr> <td>{{.ID}}</td><td>{{.Name}}</td><td>{{.Price}}</td><td>{{.Quantity}}</td>
		  </tr>
		{{end}}
	  </table>
	</body>
  </html>`
	t, err := template.New("findbyid").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	if err := t.Execute(writer, result); err != nil {
		log.Fatal(err)
	}
}

func (pc *ProductController) Update(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
	id, err := strconv.Atoi(request.FormValue("id"))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
	price, err := strconv.Atoi(request.FormValue("price_update"))
	if err != nil {
		panic(err)
	}
	quantity, err := strconv.Atoi(request.FormValue("quantity_update"))
	if err != nil {
		panic(err)
	}

	products := &entity.Product{
		ID:       int64(id),
		Name:     request.FormValue("name_update"),
		Price:    float32(price),
		Quantity: quantity,
	}

	err = pc.ProductUsecase.Update(request.Context(), products)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}
func (pc *ProductController) Delete(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
	id, err := strconv.Atoi(request.FormValue("id_delete"))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}

	err = pc.ProductUsecase.Delete(request.Context(), id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
	writer.WriteHeader(http.StatusOK)
}
