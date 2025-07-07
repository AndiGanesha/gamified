package service

import (
	"context"

	"github.com/AndiGanesha/gamified/application"
	"github.com/AndiGanesha/gamified/configuration"
	"github.com/AndiGanesha/gamified/model"
	"github.com/AndiGanesha/gamified/repository"
	"github.com/go-redis/redis/v8"
)

// define interface
type IProductService interface {
	GetProducts() ([]model.Product, error)
	GetSales(token string) ([]model.SalesTransaction, error)
	BuyProduct(token string, productId string) error
}

// define a scallable struct if needed in the future
type ProductService struct {
	config   *configuration.Configuration
	authRepo repository.IAuthenticationRepository
	redis    redis.Client
	context  context.Context
}

// create stock service func
func NewProductService(app *application.App) IProductService {
	return &ProductService{
		authRepo: repository.NewAuthenticationRepository(app),
		redis:    app.Redis,
		config:   app.Configuration,
		context:  app.Context,
	}
}
