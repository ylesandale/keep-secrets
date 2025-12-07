package main

import (
	"log"
	"net/http"
	"os"

	"keep-secrets/server/handlers"
	"keep-secrets/server/middleware"
	"keep-secrets/server/storage"

	"github.com/gorilla/mux"
)

func main() {
	db, err := storage.InitDB("secrets.db")
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	secretHandler := handlers.NewSecretHandler(db)

	r := mux.NewRouter()

	r.Use(middleware.CORS)

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/secrets", secretHandler.CreateSecret).Methods("POST", "OPTIONS")
	api.HandleFunc("/secrets/{id}", secretHandler.GetSecret).Methods("GET", "OPTIONS")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}


