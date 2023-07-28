package infrastructure

import (
	"log"
	"net/http"

	"github.com/AndiGanesha/authentication/application"
	"github.com/AndiGanesha/authentication/controller"
	"github.com/gorilla/mux"
)

func ServeHTTP(app *application.App) {
	router := mux.NewRouter().StrictSlash(true)
	setup := controller.SetupController(app)
	router.HandleFunc("sign_up", setup.AuthenticationController.SignIn).Methods("POST")
	router.HandleFunc("sign_in", setup.AuthenticationController.SignIn).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}
