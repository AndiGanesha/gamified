package controller

import (
	"encoding/json"
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
	if ok, err := c.authService.VerifyUserFromDB(user); err != nil {
		res.Error = err.Error()
		writeResponse(w, res, 500)
		return
	} else if ok {
		res.Result.Others = "username already taken"
		writeResponse(w, res, 400)
		return
	}

	// Create user in DB
	if err := c.authService.CreateUser(user); err != nil {
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

func (c *AuthenticationController) SignIn(w http.ResponseWriter, r *http.Request) {
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
