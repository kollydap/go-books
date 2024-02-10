package main

import (
	"core/controllers"
	"core/models"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	models.ConnectDatabase()
	router_group := router.Group("/books")
	{
		router_group.GET("", controllers.FindBooks)
		router_group.POST("", controllers.CreateBook)
		router_group.GET(":id", controllers.FindBook)
		router_group.PATCH(":id", controllers.UpdateBook)
		router_group.DELETE(":id", controllers.DeleteBook)
	}
	err := router.Run()
	if err != nil {
		return
	}
}
