package main

import (
	"ApiGo/controller"
	"ApiGo/db"
	"ApiGo/repository"
	"ApiGo/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	//camada de repository
	ProductRepository := repository.NewProductRepository(dbConnection)
	//camada usecase
	ProductUsecase := usecase.NewProductUsecase(ProductRepository)
	//camada de controllers
	ProductController := controller.NewProductController(ProductUsecase)

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)
	server.POST("/product", ProductController.CreateProduct)
	server.Run(":8000")
}
