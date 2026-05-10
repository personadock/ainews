package main

import (
	"log"
	"net/http"
	"os"

	"github.com/personadock/ainews/internal/site"
)

func main() {
	addr := os.Getenv("PORT")
	if addr == "" {
		addr = "8080"
	}

	logger := log.New(os.Stdout, "", log.LstdFlags)
	server, err := site.New()
	if err != nil {
		logger.Fatalf("build server: %v", err)
	}

	logger.Printf("ainews listening on :%s", addr)
	if err := http.ListenAndServe(":"+addr, server); err != nil {
		logger.Fatalf("serve: %v", err)
	}
}
