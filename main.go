package main

import (
	"github.com/KarMint26/gin-restapi/controllers/productcontoller/productcontroller"
	"github.com/KarMint26/gin-restapi/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/products", productcontroller.Index)
	r.GET("/api/product/:id", productcontroller.Show)
	r.POST("/api/product", productcontroller.Create)
	r.PUT("/api/product/:id", productcontroller.Update)
	r.DELETE("/api/products/:id", productcontroller.Delete)

	r.Run(":8001")
}
