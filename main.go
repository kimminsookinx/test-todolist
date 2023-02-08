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
	"net/http"
	"os"

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

	//TODO: delete test ping url
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	//TODO: seperate routes for clean main func
	todoRoute := r.Group("/todo")
	{
		todo := new(controllers.TodoController)

		todoRoute.GET("/list", todo.GetList)
		todoRoute.POST("", todo.PostItem)
		todoRoute.PUT("/:todoItemId/done", todo.UpdateDoneFlag) //RESTful -> REST : PUT -> PATCH, idempotency?
		todoRoute.PUT("/:todoItemId/desc", todo.UpdateDesc)     //RESTful -> REST : PUT -> PATCH, idempotency?
	}
	r.Run(":" + os.Getenv("TODO_APP_PORT"))
}
