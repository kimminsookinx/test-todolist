package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kimminsookinx/test-todolist/forms"
	"github.com/kimminsookinx/test-todolist/models"
)

type TodoController struct{}

var todoItemModel = new(models.TodoItemModel)
var todoItemForm = new(forms.TodoItemForm)

func (ctrl TodoController) GetList(c *gin.Context) {
	data, err := todoItemModel.TodoItemList()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Message": "list failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

func (ctrl TodoController) PostItem(c *gin.Context) {

	var form forms.CreateTodoItemForm

	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
		message := todoItemForm.PostTodoItem(validationErr)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
		return
	}

	todoItemId, err := todoItemModel.Post(form)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Article could not be created"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "todoItem created", "id": todoItemId})
}
