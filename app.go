package main

import (
	"Timos-API/Projects/database"
	"Timos-API/Projects/projects"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	database.Connect()

	router := mux.NewRouter()
	router.Use(routerMw)
	router.StrictSlash(true)

	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"Authorization", "Content-Type", "Origin"},
		AllowedMethods: []string{"POST", "GET", "DELETE", "PATCH", "OPTIONS"},
	}).Handler(router)

	projects.RegisterRoutes(router)

	server := &http.Server{
		Addr:         os.ExpandEnv("${host}:3000"),
		Handler:      handler,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	defer log.Fatal(server.ListenAndServe())

}

func routerMw(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		next.ServeHTTP(w, r)
	})
}
