package main

import (
	"auth-service/internal/config"
	"auth-service/internal/dependencies"
	"auth-service/internal/gen/auth/v1/v1connect"
	"auth-service/internal/server"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	load := config.Load()
	deps, err := dependencies.Initialize(load)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	authServer := server.NewAuthServer(deps.AuthService)
	mux := http.NewServeMux()
	path, handler := v1connect.NewAuthServiceHandler(authServer)
	mux.Handle(path, handler)
	go func() {
		log.Println("Starting server on localhost:8080")
		if err := http.ListenAndServe("localhost:8080", h2c.NewHandler(mux, &http2.Server{})); err != nil {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Create a channel to listen for OS signals.
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	// Block until a signal is received.
	sig := <-sigCh
	log.Printf("Received signal: %v", sig)

	// Shutdown the server gracefully.
	if err := deps.ShutDown(); err != nil {
		log.Fatalf("Shutdown error: %v", err)
	}
	log.Println("Server stopped gracefully")
}
