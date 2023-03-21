/*
	references
	https://blog.logrocket.com/building-microservices-go-gin/
	https://github.com/Massad/gin-boilerplate
	https://blog.techchee.com/build-a-rest-api-with-golang-gin-and-mysql/
*/

package main

import (
	"github.com/kimminsookinx/test-todolist/cmd/todo-list/app"
)

//TODO: add SSL

func main() {
	app.Run()
}
