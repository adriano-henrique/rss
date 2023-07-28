package main

import (
	"log"
	"net/http"
	"os"

	api "github.com/adriano-henrique/rss/api"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func mainRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/healthz", api.ReadinessHandler)
	r.Get("/err", api.ErrorHandler)
	return r
}

func main() {
	godotenv.Load(".env")
	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Fatal("PORT environment variable is not set")
	}

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	router.Mount("/v1", mainRouter())

	newServer := &http.Server{
		Addr:    ":" + PORT,
		Handler: router,
	}

	log.Printf("Serving on port: %s\n", PORT)
	log.Fatal(newServer.ListenAndServe())
}
