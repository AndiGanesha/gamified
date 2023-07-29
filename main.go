package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/AndiGanesha/authentication/application"
	"github.com/AndiGanesha/authentication/infrastructure"
)

func check(err error) {
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

func cancelOnInterrupt(app *application.App, srv io.Closer) {
	c := make(chan os.Signal, 1)
	signal.Notify(c,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	go func() {
		fmt.Printf("system call: %+v", <-c)
		srv.Close()
		app.ContextCancel()
	}()
}

func run(app *application.App) io.Closer {
	server, err := infrastructure.ServeHTTP(app)
	check(err)
	<-app.Context.Done()
	return server
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// create context with cancel() callback function
	ctx, cancel := context.WithCancel(context.Background())
	// intialize app
	app, err := application.NewApp(ctx, cancel)
	check(err)
	defer app.Close()

	// interruptable apps
	server := run(app)
	cancelOnInterrupt(app, server)

}
