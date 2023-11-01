package productcontroller

import (
	"net/http"

	"github.com/KarMint26/gin-restapi/models"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var products []models.Product

	models.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"products": products})
}

func Show(c *gin.Context) {

}

func Create(c *gin.Context) {

}

func Update(c *gin.Context) {

}

func Delete(c *gin.Context) {

}
