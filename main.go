package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jacky/go-crud/controllers"
	"github.com/jacky/go-crud/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/posts", controllers.PostsCreate)

	r.GET("/posts", controllers.PostIndex)

	r.Run()
}
