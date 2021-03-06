package main

import (
	"log"
	"net"
	"os"

	"github.com/breaking-fullstack/forever-server/database"
	"github.com/breaking-fullstack/forever-server/service"
	"github.com/breaking-fullstack/forever-server/verifier"
	"github.com/sethvargo/go-signalcontext"
)

func main() {
	ctx, cancel := signalcontext.OnInterrupt()
	defer cancel()

	srv := NewServer(getRunAddr(),
		service.NewMusic(database.NewInMem()),
		verifier.NewFirebase(),
	)

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
