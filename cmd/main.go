package main

import (
	"log"

	"github.com/Spudymun/todo"
	"github.com/Spudymun/todo/pkg/handler"
	"github.com/Spudymun/todo/pkg/repository"
	"github.com/Spudymun/todo/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while runnng httserver: %s", err.Error())
	}
}
