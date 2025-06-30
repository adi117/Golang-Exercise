package http

import "github.com/gofiber/fiber/v2"

type Router struct {
	App               *fiber.App
	ProductController *ProductController
}

type router interface {
	Setup()
	registerPublicEndpoints()
}

func NewRouter(app *fiber.App, productController *ProductController) router {
	return &Router{
		App:               app,
		ProductController: productController,
	}
}

func (r *Router) Setup() {
	r.registerPublicEndpoints()
}

func (r *Router) registerPublicEndpoints() {
	r.App.Post("/products", r.ProductController.CreateProduct)
	r.App.Get("/products", r.ProductController.GetAllProducts)
	r.App.Get("/products/:id", r.ProductController.GetProductByID)
}
