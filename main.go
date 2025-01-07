package main

import (
	"log"
	"net/http"

	"github.com/gasBlar/GoGoManager/api/v1/routes"
	"github.com/gasBlar/GoGoManager/config"
	"github.com/gasBlar/GoGoManager/db"
)

func main() {
	config.InitEnv()
	db, err := db.Init()
	if err != nil {
		log.Fatalf("Error initializing the database: %v", err)
	}
	defer db.Close()

	r := routes.InitRoutes()

	port := config.GetEnv("APP_PORT")
	log.Println("Starting server on :" + port + "...")
	if err := http.ListenAndServe("localhost:"+port+"", r); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
