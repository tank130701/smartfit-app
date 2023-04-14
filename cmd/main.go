package main

import (
	"log"
	myapp "my-app"
	"my-app/pkg/handler"
	postgres "my-app/pkg/postgres"
	"my-app/pkg/repository"
)

func main() {
	db, err := postgres.NewPostrgesConnection("178.20.41.74", "root", "root", 5432, "train-app")
	if err != nil {
		log.Fatal(err)
	}
	
	repos := repository.NewRepository(db)
	handlers := handler.NewHandler(repos)
	srv := new(myapp.Server)

	err = srv.Run("8080", handlers.InitRoutes())
	if err != nil {
		log.Fatal(err)
	}
}
