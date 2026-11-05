package main

import (
	"log"
	"net/http"
	"time"

	"github.com/RR190701/ecomGo/internal/products"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config
}

func (app *application) Server() http.Handler {
	//http hnadler is package agnostic
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.ClientIPFromRemoteAddr) // pick one ClientIPFrom* based on your infra, see below
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("healthy"))
	})

	listProducts := products.NewHandler(nil)
	r.Get("/products", listProducts.ListProducts)
	return r

}
func (app *application) Start(h http.Handler) error {

	server := &http.Server{
		Addr:         app.config.address,
		Handler:      h,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	log.Printf("Starting server on %v", app.config.address)
	return server.ListenAndServe()
}

type config struct {
	address string
	dbConfig
}

type dbConfig struct {
	dsn string
}
