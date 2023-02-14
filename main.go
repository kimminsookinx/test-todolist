/*
	references
	https://blog.logrocket.com/building-microservices-go-gin/
	https://github.com/Massad/gin-boilerplate
	https://blog.techchee.com/build-a-rest-api-with-golang-gin-and-mysql/

	TODO: get linter
*/

package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/kimminsookinx/test-todolist/db"
	"github.com/kimminsookinx/test-todolist/router"
)

func main() {
	//DESC: load environment variables
	err := godotenv.Load("todo.env")
	if err != nil {
		log.Fatal("error: failed to load env")
		return
	}

	//DEC: initialize DB connection
	db.Init()

	//DESC: initialize and run gin server
	router.Init()
}
