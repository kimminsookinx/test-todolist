/*
	references
	https://blog.logrocket.com/building-microservices-go-gin/
	https://github.com/Massad/gin-boilerplate
	https://blog.techchee.com/build-a-rest-api-with-golang-gin-and-mysql/
*/

package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kimminsookinx/test-todolist/controllers"
	"github.com/kimminsookinx/test-todolist/db"
)

func main() {
	err := godotenv.Load("todo.env")
	if err != nil {
		log.Fatal("error: failed to load env")
	}

	r := gin.Default()
	db.Init()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	todoRoute := r.Group("/todo")
	{
		todo := new(controllers.TodoController)

		todoRoute.GET("/list", todo.GetList)
		todoRoute.POST("", todo.GetList)
	}
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
