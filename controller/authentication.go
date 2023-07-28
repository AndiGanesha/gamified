package controller

import (
	"net/http"

	"github.com/AndiGanesha/authentication/application"
	"github.com/AndiGanesha/authentication/model"
	"github.com/AndiGanesha/authentication/service"
)

// define interface
type IAuthenticationController interface {
	SignUp(w http.ResponseWriter, r *http.Request)
	SignIn(w http.ResponseWriter, r *http.Request)
}

// define a scallable struct if needed in the future
type AuthenticationController struct {
	authService service.IAuthService
}

// create stock service func
func NewAuthenticationController(app *application.App) IAuthenticationController {
	return &AuthenticationController{
		authService: service.NewAuthenticationService(app),
	}
}

func (c *AuthenticationController) SignUp(w http.ResponseWriter, r *http.Request) {
	c.authService.SignUp(model.User{})
}

func (c *AuthenticationController) SignIn(w http.ResponseWriter, r *http.Request) {
	c.authService.SignUp(model.User{})
}
