// @title Auth Service
// @version 0.1
// @description Authentication service
package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/maykonlf/go-auth-service/internal/handlers"
	"log"
	"net/http"
)

func main()  {
	router := chi.NewRouter()

	router.Use(middleware.DefaultCompress)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/healthz", handlers.LivenessHandler)
	router.Route("/login", handlers.NewLoginHandler().Router)

	log.Fatal(http.ListenAndServe(":3000", router))
}
