package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kimminsookinx/test-todolist/controllers"
)

func createRoute() {
	//TODO: delete test ping url
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	//TODO: add delete url
	//TODO: create put for compatability
	todoRoute := r.Group("/todo")
	{
		todo := new(controllers.TodoController)

		todoRoute.GET("/list", todo.GetList)
		todoRoute.POST("", todo.PostItem)
		todoRoute.PUT("/:todoItemId/done", todo.UpdateDoneFlag) //RESTful -> REST : PUT -> PATCH, idempotency?
		todoRoute.PATCH("/:todoItemId/desc", todo.UpdateDesc)
	}
}
