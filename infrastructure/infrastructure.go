package infrastructure

import (
	"io"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/AndiGanesha/authentication/application"
	"github.com/AndiGanesha/authentication/controller"
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
	router.HandleFunc("/sign_up", setup.AuthenticationController.SignUp).Methods("POST")
	router.HandleFunc("/sign_in", setup.AuthenticationController.SignIn).Methods("POST")

	return router
}

func (ln tcpKeepAliveListener) Accept() (c net.Conn, err error) {
	tc, err := ln.AcceptTCP()
	if err != nil {
		return
	}
	tc.SetKeepAlive(true)
	tc.SetKeepAlivePeriod(3 * time.Minute)
	return tc, nil
}
