package http

import (
	"github.com/adi117/Golang-Exercise/internal/model"
	"github.com/adi117/Golang-Exercise/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ProductController struct {
	Usecase usecase.ProductUsecase
	Log     *logrus.Logger
}

func NewProductController(uc *usecase.ProductUsecase, log *logrus.Logger) *ProductController {
	return &ProductController{
		Usecase: *uc,
		Log:     log,
	}
}

func (p *ProductController) CreateProduct(ctx *fiber.Ctx) error {
	request := new(model.CreateProductRequest)
	if err := ctx.BodyParser(request); err != nil {
		p.Log.WithError(err).Error("failed to parse request body")
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[any]{
			Message: "invalid request body format",
			Success: false,
		})
	}

	resp, err := p.Usecase.CreateProduct(ctx.Context(), request)
	if err != nil {
		p.Log.WithError(err).Error("failed to create product")
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.WebResponse[*model.CreateProductResponse]{
			Data:    resp,
			Success: false,
			Message: "failed to create product",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(model.WebResponse[*model.CreateProductResponse]{
		Data:    resp,
		Success: true,
		Message: "product created successfully",
	})
}

func (p *ProductController) GetAllProducts(ctx *fiber.Ctx) error {
	products, err := p.Usecase.GetAllProducts(ctx.Context())
	if err != nil {
		p.Log.Error("failed to fetch products")
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.WebResponse[any]{
			Success: false,
			Message: "failed to fetch products",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(model.WebResponse[any]{
		Data:    products,
		Success: true,
		Message: "products fetched successfully",
	})
}
