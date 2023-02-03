package main

import (
	"Gin/initializers"
	"Gin/models"
)

func init() {

	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}