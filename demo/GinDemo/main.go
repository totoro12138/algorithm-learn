package main

import (
	"algorithm-learn/demo/GinDemo/run"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "gin_hello_world_3!")
	})

	run.GinRun(router, ":8080")
}
