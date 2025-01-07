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
	database := db.InitDb()
	defer database.Close()

	r := routes.InitRoutes()

	port := config.GetEnv("APP_PORT")
	log.Println("Starting server on :" + port + "...")
	if err := http.ListenAndServe(":"+port+"", r); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
