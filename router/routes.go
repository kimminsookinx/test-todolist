package router

import "github.com/kimminsookinx/test-todolist/controllers"

/*
	API endpoint naming conventions (clusterapi -> k8s)
	https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#naming-conventions
*/

func createRoute() {
	//TODO: add delete url
	//TODO: create put for compatability?
	todoRoute := r.Group("/v1/todos")
	{
		todoController := new(controllers.TodoController)

		todoRoute.GET("", todo.GetList)
		todoRoute.POST("", todo.PostItem)
		todoRoute.PUT("/:todoItemId/done", todo.UpdateDoneFlag) //RESTful -> REST : PUT -> PATCH, idempotency?
		todoRoute.PATCH("/:todoItemId/desc", todo.UpdateDesc)
		//todoRoute.DELETE("/:todoItemID",todo.Delete)
		//TODO: only way to init runtime variables? (env vars set at runtime)
		todoController.Init()
	}
}
