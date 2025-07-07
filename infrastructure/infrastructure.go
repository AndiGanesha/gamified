package infrastructure

import (
	"io"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/AndiGanesha/gamified/application"
	"github.com/AndiGanesha/gamified/controller"
	"github.com/gorilla/mux"
)

type tcpKeepAliveListener struct {
	*net.TCPListener
}

func ServeHTTP(app *application.App) (sc io.Closer, err error) {
	var (
		listener net.Listener
	)
	srv := &http.Server{
		Addr:    app.Configuration.Server.HTTP,
		Handler: handler(app),
	}
	listener, err = net.Listen("tcp", app.Configuration.Server.HTTP)
	if err != nil {
		return nil, err
	}

	go func() {
		log.Println("serve HTTP for", app.Configuration.Server.HTTP)
		err := srv.Serve(tcpKeepAliveListener{listener.(*net.TCPListener)})
		if err != nil {
			log.Println("HTTP Server Error - ", err)
		}
	}()

	return listener, nil
}

func handler(app *application.App) http.Handler {
	router := mux.NewRouter()

	// make it easier to add more handler
	setup := controller.SetupController(app)
	router.HandleFunc("/register", setup.AuthenticationController.SignUp).Methods("POST")
	router.HandleFunc("/login", setup.AuthenticationController.SignIn).Methods("POST")
	router.HandleFunc("/buy-product", setup.ProductController.BuyProduct).Methods("POST")
	router.HandleFunc("/get-products", setup.ProductController.GetProducts).Methods("GET")
	router.HandleFunc("/get-transactions", setup.ProductController.GetTransactions).Methods("POST")

	return router
}

func (ln tcpKeepAliveListener) Accept() (c net.Conn, err error) {
	tc, err := ln.AcceptTCP()
	if err != nil {
		return
	}
	err = tc.SetKeepAlive(true)
	if err != nil {
		return
	}
	err = tc.SetKeepAlivePeriod(3 * time.Minute)
	if err != nil {
		return
	}
	return tc, nil
}
