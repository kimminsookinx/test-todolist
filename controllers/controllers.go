// TODO: implement common error handler(middleware?)
// TODO: gorp -> gorm
// TODO: unit test
// TODO: logs
package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kimminsookinx/test-todolist/forms"
	"github.com/kimminsookinx/test-todolist/models"
)

/*
	HTTP status code references
		HTTP methods
		https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#verbs-on-resources
		HTTP response code
		https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#http-status-codes
*/

/*
	HTTP status code responses for todo-app methods

		GET
			-200	get success, reponse body has json
			-400	wrong query params or type
			-500	internal server error
		POST
			-201	item creation success
			-400	wrong request body json
			-500	internal server error
		PATCH
			-200	item update success
			-201	item creation success -> not used
			-400	wrong request body json
			-500	internal server error
		DELETE
			-204	item deletion success
			-400	already deleted item at id
			-404	id not found ()
			-500	internal server error
		PUT - not in use
			-200
			-201
			-400
			-500
*/

type TodoController struct{}

var todoItemModel = new(models.TodoItemModel)
var todoItemForm = new(forms.TodoItemForm)

func (ctrl TodoController) Init() {
	todoItemModel.Init()
}

func (ctrl TodoController) GetList(c *gin.Context) {
	var data []models.TodoItem
	var err error

	//NOTE: probably not clean (ex: multiple query param)
	//		maybe use interceptors(does this even exist)? -> see middleware (https://stackoverflow.com/questions/69948784/how-to-handle-errors-in-gin-middleware)

	var formString forms.GetQueryFormString

	if validationErr := c.Bind(&formString); validationErr != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	data, err = todoItemModel.SelectTodoItem(c.Request.URL.Query())

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

func (ctrl TodoController) PostItem(c *gin.Context) {

	var form forms.CreateTodoItemForm

	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
		// todoItemForm.CheckDesc(validationErr)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	//TODO: select between empty body or return last insert id
	//todoItemId, err := todoItemModel.InsertTodoItem(form)
	_, err := todoItemModel.InsertTodoItem(form)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	//c.JSON(http.StatusCreated, gin.H{"id": todoItemId})
	c.Status(http.StatusCreated)

}

func (ctrl TodoController) PatchItemDoneFlag(c *gin.Context) {
	//new resource creation not allowed
	todoItemId, err := strconv.ParseInt(c.Param("todoItemId"), 10, 64)
	if todoItemId == 0 || err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var form forms.UpdateDoneTodoItemForm
	if validationErr := c.BindJSON(&form); validationErr != nil {
		todoItemForm.CheckDoneFlag(validationErr)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = todoItemModel.UpdateTodoItemSetDoneById(todoItemId, form)
	if err != nil {
		//sql update error
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Message": "todoitem could not be updated"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "todo updated"})
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
		message := todoItemForm.CheckDesc(validationErr)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
		return
	}

	err = todoItemModel.UpdateTodoItemSetDescById(todoItemId, form)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"Message": "todoitem could not be updated"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "todo updated"})
}

//NOTE: not using db instance's time might introduce time discrepancy
//		but this exercise is getting used to go, so we insert datetime from this todo-app
//		the better way should be using triggers to update if deleted column becomes true
//		(https://stackoverflow.com/questions/37856582/on-update-current-timestamp-for-only-one-column-in-mysql)
//TODO: probably should be transaction

func (ctrl TodoController) DeleteItem(c *gin.Context) {
	//check row existence
	idString := c.Param("todoItemId")
	todoItemId, err := strconv.ParseInt(idString, 10, 64)
	if todoItemId == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": err.Error()})
		return
	}

	exists, err := todoItemModel.CheckRowExistenceByIdAndDeleted(todoItemId, false)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	if !exists {
		//no row exists, return 404
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	success, err := todoItemModel.UpdateTodoItemSetDeletedIsFalseById(todoItemId)
	if !success {
		//TODO: do not compare error.Error() string, use wrapper
		if err.Error() == "updated 0 records" {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		return
	}

	c.Status(http.StatusNoContent)
}
