package router

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func Init() {
	r = gin.Default()
	createRoute()

	//NOTE: Default 404
	r.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})
}
func Run() {
	r.Run(":" + os.Getenv("TODO_APP_PORT"))
}
