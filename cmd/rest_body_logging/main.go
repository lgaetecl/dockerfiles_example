package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/lgaetecl/restbodylogging/internal/handler"
)

var (
	port = 3000
)

func main() {
	r := chi.NewRouter()

	r.Post("/", handler.BodyLogging)
	r.Get("/", handler.BodyLogging)

	log.Printf("Server Listen on: %d \n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}
