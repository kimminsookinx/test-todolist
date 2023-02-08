package controllers

import (
	"fmt"
	"net/http"
	"strconv"

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
		message := todoItemForm.CheckDesc(validationErr)
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

func (ctrl TodoController) UpdateDoneFlag(c *gin.Context) {
	idString := c.Param("todoItemId")

	todoItemId, err := strconv.ParseInt(idString, 10, 64)
	if todoItemId == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter"})
		return
	}

	var form forms.UpdateDoneTodoItemForm
	if validationErr := c.BindJSON(&form); validationErr != nil {
		fmt.Print(validationErr.Error())
		message := todoItemForm.CheckDoneFlag(validationErr)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
		return
	}

	err = todoItemModel.UpdateDone(todoItemId, form)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Message": "todoitem could not be updated"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "todo updated"})
}

func (ctrl TodoController) UpdateDesc(c *gin.Context) {
	idString := c.Param("todoItemId")

	todoItemId, err := strconv.ParseInt(idString, 10, 64)
	if todoItemId == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter"})
		return
	}

	var form forms.UpdateDescTodoItemForm
	if validationErr := c.BindJSON(&form); validationErr != nil {
		fmt.Print(validationErr.Error())
		message := todoItemForm.CheckDesc(validationErr)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
		return
	}

	err = todoItemModel.UpdateDesc(todoItemId, form)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Message": "todoitem could not be updated"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "todo updated"})
}
