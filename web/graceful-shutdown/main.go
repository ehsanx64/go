package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	// Configure the web server
	srv := &http.Server{
		Addr: ":8000",
	}

	go func() {
		// We don't want http.ErrServerClosed to be logged as an error
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	// Define a os.Signal channel to be able to block the application until
	// receive a signal from OS
	quit := make(chan os.Signal, 1)

	// Send a signal to quit channel when program is terminated (Ctrl-C etc.)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// Wait no more than 30 seconds for any request
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Finally shutdown the server in a graceful manner
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
