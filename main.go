package main

import (
	// core packages
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"

	// internal packages
	"github.com/ctit-team/oid-registry/config"

	// external packages
	"github.com/ctit-team/nestederror"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	// setup routes
	mux := http.NewServeMux()

	// setup http server
	serv := http.Server{
		Addr:    config.Main.HTTP.Listener.Address,
		Handler: mux,
	}

	noti := make(chan os.Signal, 1)
	signal.Notify(noti, os.Interrupt)

	var tasks sync.WaitGroup
	tasks.Add(1)

	go func() {
		select {
		case <-noti:
			serv.Shutdown(context.Background())
		}
		tasks.Done()
	}()

	// run http server
	var err error

	if err = serv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		err = nestederror.New(err, "client serving error")
	} else {
		err = nil
	}

	// clean up
	signal.Stop(noti)
	close(noti)
	tasks.Wait()

	return err
}
