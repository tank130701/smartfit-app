package main

import (
	myapp "my-app"
	"my-app/internal/handler"
	"my-app/internal/repository"
	"my-app/internal/services"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	db, err := repository.NewPostgresConnection("194.87.110.172", "val1", "root", 5432, "my-app")
	if err != nil {
		logrus.Fatal(err)
	}

	repos := repository.NewRepository(db)
	services := services.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(myapp.Server)

	err = srv.Run("8080", handlers.InitRoutes())
	if err != nil {
		logrus.Fatal(err)
	}
}
