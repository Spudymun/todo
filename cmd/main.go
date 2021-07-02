package main

import (
	"log"

	"github.com/Spudymun/todo"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todo.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error occured while runnng httserver: %s", err.Error())
	}
}
