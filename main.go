package main

import (
	"log"
	"net/http"

	"github.com/gasBlar/GoGoManager/api/v1/routes"
	"github.com/gasBlar/GoGoManager/db"
)

func main() {
	db, err := db.Init()
	if err != nil {
		log.Fatalf("Error initializing the database: %v", err)
	}
	defer db.Close()

	routes.SetupRoutes()

	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
