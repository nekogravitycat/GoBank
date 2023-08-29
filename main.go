package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", hello)
	r.Run() // listen and serve on 0.0.0.0:8080
}

func hello(c *gin.Context) {
	msg := []byte("hello world")
	c.Data(http.StatusOK, "text/plain; charset=utf-8", msg)
}
