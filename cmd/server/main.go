package main

import (
	"log"

	"github.com/nomad31415/goprolog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8084")
	log.Fatal(srv.ListenAndServe())
}
