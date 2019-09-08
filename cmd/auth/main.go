// Schemes: http
// Host: http://auth-service
// BasePath: /v1
// Version: 0.0.1
//
// Consumes:
// - application/json
// - application/xml
//
// Produces:
// - application/json
// - application/xml
//
//
// swagger:meta
package main

import (
	"github.com/go-chi/chi"
	"github.com/maykonlf/go-auth-service/internal/handlers"
	"log"
	"net/http"
)

func main()  {
	router := chi.NewRouter()

	router.Get("/healthz", handlers.LivenessHandler)

	log.Fatal(http.ListenAndServe(":3000", router))
}
