package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/AndiGanesha/gamified/application"
)

type Controllers struct {
	AuthenticationController IAuthenticationController
	ProductController        IProductController
}

// helper to setup if there will be many consumer in the future
func SetupController(app *application.App) *Controllers {
	return &Controllers{
		AuthenticationController: NewAuthenticationController(app),
		ProductController:        NewProductController(app),
	}
}

// helper to write response
func writeResponse(w http.ResponseWriter, body interface{}, status int) {
	resp, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	_, err = w.Write(resp)
	if err != nil {
		log.Fatal(err)
	}
}
