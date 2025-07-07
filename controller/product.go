package controller

import (
	"encoding/json"
	"net/http"

	"github.com/AndiGanesha/gamified/application"
	"github.com/AndiGanesha/gamified/model"
	"github.com/AndiGanesha/gamified/service"
)

// define interface
type IProductController interface {
	GetProducts(w http.ResponseWriter, r *http.Request)
	GetTransactions(w http.ResponseWriter, r *http.Request)
	BuyProduct(w http.ResponseWriter, r *http.Request)
}

// define a scallable struct if needed in the future
type ProductController struct {
	productService service.IProductService
}

// create product controller func
func NewProductController(app *application.App) IProductController {
	return &ProductController{
		productService: service.NewProductService(app),
	}
}

func (c *ProductController) GetProducts(w http.ResponseWriter, r *http.Request) {
	//Read input from request
	var (
		product model.Product
		res     model.ResponseSign
	)

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		res.Error = err.Error()
		writeResponse(w, res, 400)
		return
	}

	// set token in redis
	if err := c.productService.g(); err != nil {
		res.Error = err.Error()
		writeResponse(w, res, 500)
		return
	}

	res.Result.Token = token
	writeResponse(w, res, 200)
}

func (c *ProductController) GetTransactions(w http.ResponseWriter, r *http.Request) {
	//Read input from request
	var (
		product model.Product
		res     model.ResponseProduct
	)

	token := r.Header.Get("Authentication")
	if token == "" {
		res.Error = "Please Login"
		writeResponse(w, res, 400)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		res.Error = err.Error()
		writeResponse(w, res, 400)
		return
	}

	// Verify username in DB
	if ok, err := c.authService.VerifyUserFromDB(user); err != nil || !ok {
		res.Error = err.Error()
		writeResponse(w, res, 500)
		return
	}

	res.Result.Token = token
	writeResponse(w, res, 200)
}

func (c *ProductController) BuyProduct(w http.ResponseWriter, r *http.Request) {
	//Read input from request
	var (
		user model.User
		res  model.ResponseSign
	)

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		res.Error = err.Error()
		writeResponse(w, res, 400)
		return
	}

	// Verify username in DB
	if ok, err := c.authService.VerifyUserFromDB(user); err != nil || !ok {
		res.Error = err.Error()
		writeResponse(w, res, 500)
		return
	}

	// Generate token
	token, err := c.authService.GenerateToken(user.Username, user.Password)
	if err != nil {
		res.Error = err.Error()
		writeResponse(w, res, 500)
		return
	}

	// set token in redis
	if err := c.authService.SetRedisToken(token, user); err != nil {
		res.Error = err.Error()
		writeResponse(w, res, 500)
		return
	}

	res.Result.Token = token
	writeResponse(w, res, 200)
}
