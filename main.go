package main

import (
	"awesomeProject/database"
	"awesomeProject/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	r := gin.Default()
	r.MaxMultipartMemory = 100 << 20
	usersGroup := r.Group("/users")
	{
		usersGroup.GET("/", routes.GetAllUsersHandler)
		usersGroup.POST("/", routes.CreateUserHandler)
		usersGroup.GET("/:id", routes.GetOneUserHandler)
	}
	blogsGroup := r.Group("/blogs")
	{
		blogsGroup.POST("/", routes.CreateBlogHandler)
		blogsGroup.POST("/:id", routes.UpdateBlogHandler)
		blogsGroup.GET("/:id", routes.GetOneBlogHandler)
		blogsGroup.POST("/upload", routes.UploadFileHandler)
	}

	r.Run(":8080")
}
