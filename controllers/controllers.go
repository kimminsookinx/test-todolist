package controllers

import (
	"net/http"

	//"strconv"
	"github.com/gin-gonic/gin"
	"github.com/kimminsookinx/test-todolist/models"
)

type TodoController struct{}

var todoItemModel = new(models.TodoItemModel)

func (ctrl TodoController) PostItem(c *gin.Context) {
	//todoDescription :=
}

func (ctrl TodoController) GetList(c *gin.Context) {
	data, err := todoItemModel.TodoItemList()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Message": "list failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}
