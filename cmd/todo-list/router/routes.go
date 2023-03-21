package router

import "github.com/kimminsookinx/test-todolist/cmd/todo-list/controllers"

/*
	API endpoint naming conventions (clusterapi -> k8s)
	https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#naming-conventions
*/

func createRoute() {
	todoRoute := r.Group("/v2/todos")
	{
		todoController := new(controllers.TodoController)

		//TODO: only way to init runtime variables? (env vars set at runtime)
		todoController.Init()

		todoRoute.GET("", todoController.GetList)
		todoRoute.POST("", todoController.PostItem)
		todoRoute.PATCH("/:todoItemId/done", todoController.PatchItemDoneFlag)
		todoRoute.PATCH("/:todoItemId/desc", todoController.UpdateDesc)
		todoRoute.DELETE("/:todoItemId", todoController.DeleteItem)
	}
}
