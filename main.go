package main

import (
	"net/http"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(requestid.New())

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "id:"+requestid.Get(c))
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
