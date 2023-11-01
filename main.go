package main

import (
	"github.com/KarMint26/gin-restapi/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/products", ProductController.Index)
	r.GET("/api/product/:id", ProductController.Show)
	r.POST("/api/product", ProductController.Create)
	r.PUT("/api/product/:id", ProductController.Update)
	r.DELETE("/api/products/:id", ProductController.Delete)

	r.Run(":8001")
}
