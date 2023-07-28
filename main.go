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
)

const (
	// filepath defined here
	filePath = "src/rawdata"
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

	// setup controller consumer
	ctrl := controller.SetupController(app)

	// interruptable apps
	cancelOnInterrupt(app)

	// declare and initiate consumer  that will be needed for calculation
	infrastructure.ConsumeKafkaMessage(app, ctrl)

	// intiate server
	infrastructure.StartHTTPServer(app, ctrl)

	<-ctx.Done()
}
