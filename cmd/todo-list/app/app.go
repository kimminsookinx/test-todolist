package app

import (
	"log"

	"github.com/joho/godotenv"

	db "github.com/kimminsookinx/test-todolist/cmd/todo-list/repository"
	"github.com/kimminsookinx/test-todolist/cmd/todo-list/router"
)

//TODO: add SSL

func Run() {
	//NOTE: load environment variables
	err := godotenv.Load("todo.env")
	if err != nil {
		log.Fatal("error: failed to load env")
		return
	}

	//NOTE: initialize DB connection
	db.Init()

	//NOTE: initialize and run gin server
	router.Init()
	router.Run()
}
