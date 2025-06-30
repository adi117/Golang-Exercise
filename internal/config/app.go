package config

import (
	"github.com/adi117/Golang-Exercise/internal/delivery/http"
	"github.com/adi117/Golang-Exercise/internal/repository"
	"github.com/adi117/Golang-Exercise/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type AppConfig struct {
	DB     *gorm.DB
	App    *fiber.App
	Log    *logrus.Logger
	Config *viper.Viper
}

func (cfg *AppConfig) Run() {
	productRepository := repository.NewProductRepository(cfg.Log, cfg.DB)

	productUseCase := usecase.NewProductUsecase(productRepository, cfg.Log, cfg.DB)

	productController := http.NewProductController(&productUseCase, cfg.Log)

	routeConfig := http.Router{
		App:               cfg.App,
		ProductController: productController,
	}

	routeConfig.Setup()
}
