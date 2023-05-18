package router

import (
	"github.com/dihanto/crud-web/controller"
	"github.com/dihanto/crud-web/middleware"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(pc *controller.ProductController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/products", middleware.Logger(pc.GetAll))
	router.GET("/products/detail/:id", middleware.Logger(pc.FindById))
	router.GET("/products/edit/:id", middleware.Logger(pc.ProductEdit))
	router.PUT("/products/:id", middleware.Logger(pc.Update))
	router.DELETE("/products/:id", middleware.Logger(pc.Delete))
	router.GET("/products/add", middleware.Logger(pc.ProductAdd))
	router.POST("/products/add", middleware.Logger(pc.Create))

	return router
}
