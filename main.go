/*
	references
	https://blog.logrocket.com/building-microservices-go-gin/
	https://github.com/Massad/gin-boilerplate
	https://blog.techchee.com/build-a-rest-api-with-golang-gin-and-mysql/
*/

package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/kimminsookinx/test-todolist/db"
	"github.com/kimminsookinx/test-todolist/router"
)

//TODO: add SSL

func main() {
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
