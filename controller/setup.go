package controller

import "github.com/AndiGanesha/authentication/application"

type Controllers struct {
	AuthenticationController IAuthenticationController
}

// setup if there will be many consumer in the future
func SetupController(app *application.App) *Controllers {
	return &Controllers{
		AuthenticationController: NewAuthenticationController(app),
	}
}
