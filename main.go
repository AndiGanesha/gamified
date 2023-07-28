package main

import (
	"context"
	"fmt"
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

func cancelOnInterrupt(app *application.App) {
	c := make(chan os.Signal, 1)
	signal.Notify(c,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	go func() {
		fmt.Printf("system call: %+v", <-c)
		app.ContextCancel()
	}()
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
	cancelOnInterrupt(app)

	// intiate server
	infrastructure.ServeHTTP(app)

	<-ctx.Done()
}
