package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"vanilla-go-api/server"
)

func main() {

	ctx := context.Background()

	serverDoneChan := make(chan os.Signal, 1)
	signal.Notify(serverDoneChan, os.Interrupt, syscall.SIGTERM)

	server := server.NewServer(":3000")

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()

	log.Println("Server started")

	<-serverDoneChan

	server.Shutdown(ctx)
	log.Println("Server stopped")
}
