package main

import (
	"log"

	todo "github.com/airvt1x/simple-todo-app"
	"github.com/airvt1x/simple-todo-app/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error while running http server: %s", err.Error())
	}
}
