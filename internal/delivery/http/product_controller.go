package http

import (
	"math"
	"strconv"

	"github.com/adi117/Golang-Exercise/internal/entity"
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

	page, _ := strconv.Atoi(ctx.Query("page", "1"))
	limit, _ := strconv.Atoi(ctx.Query("limit", "10"))

	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 1
	}

	offset := (page - 1) * limit

	products, total, err := p.Usecase.GetAllProducts(ctx.Context(), limit, offset)
	if err != nil {
		p.Log.Error("failed to fetch products")
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.WebResponse[any]{
			Success: false,
			Message: "failed to fetch products",
		})
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	paginated := model.PaginatedResponse[*entity.Product]{
		Items:       products,
		Page:        page,
		Limit:       limit,
		Total:       total,
		TotalPages:  totalPages,
		HasNext:     page < totalPages,
		HasPrevious: page > 1,
	}

	return ctx.Status(fiber.StatusOK).JSON(model.WebResponse[model.PaginatedResponse[*entity.Product]]{
		Data:    paginated,
		Success: true,
		Message: "products fetched successfully",
	})
}

func (p *ProductController) GetProductByID(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")

	if err != nil {
		p.Log.WithError(err).Error("failed to parse id")
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[any]{
			Message: "invalid product id format",
			Success: false,
		})
	}

	product, err := p.Usecase.GetProductByID(ctx.Context(), int64(id))

	if err != nil {
		if err.Error() == "Product not found" {
			return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[any]{
				Message: "Product not found",
				Success: false,
			})
		}

		p.Log.WithError(err).Error("failed to get product")
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.WebResponse[any]{
			Message: "failed to retrieve product details",
			Success: false,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(model.WebResponse[any]{
		Data:    product,
		Success: true,
		Message: "products fetched successfully",
	})
}

func (p *ProductController) UpdateProduct(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")

	if err != nil {
		p.Log.WithError(err).Error("failed to parse id")
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[any]{
			Message: "invalid product id format",
			Success: false,
		})
	}

	request := new(model.UpdateProductRequest)

	if err := ctx.BodyParser(request); err != nil {
		p.Log.WithError(err).Error("failed to parse request body")
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[any]{
			Message: "invalid request body format",
			Success: false,
		})
	}

	product, err := p.Usecase.UpdateProduct(ctx.Context(), int64(id), request)

	if err != nil {
		if err.Error() == "Product not found" {
			return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[any]{
				Message: "Product not found",
				Success: false,
			})
		}

		p.Log.WithError(err).Error("failed to get product")
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.WebResponse[any]{
			Message: "failed to retrieve product details",
			Success: false,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(model.WebResponse[any]{
		Data:    product,
		Success: true,
		Message: "product updated successfully",
	})
}
