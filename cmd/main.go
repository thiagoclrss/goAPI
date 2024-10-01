package main

import (
	"ApiGo/controller"
	"ApiGo/db"
	"ApiGo/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	ProductUsecase := usecase.NewProductUsecase()

	ProductController := controller.NewProductController(ProductUsecase)

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)
	server.Run(":8000")
}
