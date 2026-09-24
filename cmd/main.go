package main

import (
	"log"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.Default()

	srv := server.NewServer(logger)

	logger.Fatal(srv.Start())
}
