package router

import (
	"os"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func Init() {
	r = gin.Default()
	createRoute()
	r.Run(":" + os.Getenv("TODO_APP_PORT"))
}
