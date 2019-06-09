package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/SUMUKHA-PK/Mail-Server/routeHandlers"
	"github.com/gorilla/mux"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {

	r := mux.NewRouter()

	// r.HandleFunc("/", routeHandlers.HandlerFunc)
	*r = routeHandlers.SetupRouting(*r)

	server := &http.Server{
		Handler: r,
		Addr:    ":8080",
	}

	LogFileLocation := os.Getenv("LogFileLocation")
	if LogFileLocation != "" {
		log.SetOutput(&lumberjack.Logger{
			Filename:   LogFileLocation,
			MaxSize:    500, // megabytes
			MaxBackups: 3,
			MaxAge:     28,   //days
			Compress:   true, // disabled by default
		})
	}

	go func() {
		log.Println("Starting Server")
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	waitForShutdown(server)
}

// waitForShutdown shuts down the server on getting a ^C signal
func waitForShutdown(server *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	server.Shutdown(ctx)

	log.Println("Shutting down")
	os.Exit(0)
}
