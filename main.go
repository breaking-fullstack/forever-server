package main

import (
	"log"
	"net"
	"os"

	"github.com/sethvargo/go-signalcontext"
)

func main() {
	ctx, cancel := signalcontext.OnInterrupt()
	defer cancel()

	srv := NewServer(getRunAddr(), nil)

	go func() {
		if err := srv.Start(); err != nil {
			log.Fatal(err)
		}
	}()

	// Wait for CTRL+C
	<-ctx.Done()

	// Stop the server
	if err := srv.Stop(); err != nil {
		log.Fatal(err)
	}
}

func getRunAddr() string {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}
	return net.JoinHostPort("", port)
}
