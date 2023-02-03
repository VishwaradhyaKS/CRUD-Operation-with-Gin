package controllers

import (
	"Gin/initializers"
	"Gin/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	//Get data of requset body

	// var body struct {
	// 	Body  string
	// 	Title string
	// }
	// c.Bind(&body)
	//Create a post

	var newpost models.Post

	c.Bind(&newpost)

	post := models.Post{Title: newpost.Title, Body: newpost.Body}

	result := initializers.DB.Create(&newpost)
	if result.Error != nil {
		c.Status(400)
		return
	}

	//Return it

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {

	//Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	//Respond with them

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {

	//Get id of url

	id := c.Param("id")

	//Get the posts
	var post models.Post
	initializers.DB.First(&post, id)

	//Respond with them

	c.JSON(200, gin.H{
		"posts": post,
	})
}

func PostsUpdate(c *gin.Context) {
	//Get the id of url
	id := c.Param("id")

	//Get the data of request body

	var newpost models.Post

	c.Bind(&newpost)

	//Find the post were updating

	var post models.Post
	initializers.DB.First(&post, id)

	//update it

	initializers.DB.Model(&post).Updates(models.Post{
		Title: newpost.Title,
		Body:  newpost.Body,
	})

	//Respond with it

	c.JSON(200, gin.H{
		"posts": post,
	})
}

func PostsDelete(c *gin.Context) {
	// Get the id of url

	id := c.Param("id")

	//Delete the post

	initializers.DB.Delete(&models.Post{}, id)

	//Respond

	c.Status(200)
}
