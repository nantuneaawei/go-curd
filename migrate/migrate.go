package main

import (
	"github.com/jacky/go-crud/initializers"
	"github.com/jacky/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
