package main

import (
	"net/http"
	"time"

	"github.com/dihanto/crud-web/app"
	"github.com/dihanto/crud-web/controller"
	"github.com/dihanto/crud-web/repository"
	"github.com/dihanto/crud-web/router"
	"github.com/dihanto/crud-web/usecase"
)

func main() {

	repo := repository.NewProductRepository(app.GetConnection())

	pu := usecase.NewProductUsecase(repo, 60*time.Second)

	pc := controller.NewProductController(pu)

	r := router.NewRouter(pc)

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: r,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
