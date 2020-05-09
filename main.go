package main

import (
	"github.com/gin-gonic/gin"
	"gintest/action"
)

func main() {
	r := gin.Default()
	r.GET("/ping", action.Test)
	r.Run() // listen and serve on 0.0.0.0:8080
}
