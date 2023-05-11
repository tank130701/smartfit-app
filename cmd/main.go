package main

import (
	"log"
	myapp "my-app"
	"my-app/internal/handler"
	postgres "my-app/internal/postgres"
	"my-app/internal/repository"
)

func main() {
	db, err := postgres.NewPostrgesConnection("194.87.110.172", "val1", "root", 5432, "my-app")
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
