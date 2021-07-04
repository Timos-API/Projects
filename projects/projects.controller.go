package projects

import (
	"fmt"

	authenticator "github.com/Timos-API/Authenticator"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	fmt.Println("Projects routes registered")
	s := router.PathPrefix("/projects").Subrouter()

	s.HandleFunc("", getAllProjects).Methods("GET")
	s.HandleFunc("/{id}", getProject).Methods("GET")

	s.HandleFunc("", authenticator.AuthMiddleware(nil, []string{"admin"})).Methods("POST")
	s.HandleFunc("/{id}", authenticator.AuthMiddleware(nil, []string{"admin"})).Methods("DELETE")
	s.HandleFunc("/{id}", authenticator.AuthMiddleware(nil, []string{"admin"})).Methods("PATCH")

}
