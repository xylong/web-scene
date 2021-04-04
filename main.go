package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web/AppInit"
	"web/models"
)

func main() {
	router := gin.Default()

	v1 := router.Group("v1")
	{
		v1.Handle(http.MethodGet, "books", func(context *gin.Context) {
			books := models.Books{}
			AppInit.GetDB().Limit(10).Order("id desc").Find(&books)
			context.JSON(http.StatusOK, books)
		})
	}

	router.Run(AppInit.Address)
}
