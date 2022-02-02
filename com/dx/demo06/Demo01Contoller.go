package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()
	v1 := engine.Group("/v1")
	{
		v1.GET("/hello", hello)
		v1.GET("/hi", hi)
	}

	engine.GET("/get", func(c *gin.Context) {
		c.String(http.StatusOK, "htllo world")
	})
	engine.Run(":8080")

}

func hello(c *gin.Context) {
	c.String(http.StatusOK, "htllo")
}

func hi(c *gin.Context) {
	c.String(http.StatusOK, "hi")
}
