package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web-scene/App"
	"web-scene/App/Services"
	"web-scene/AppInit"
)

func main() {
	router := gin.Default()

	v1 := router.Group("v1")
	{
		bookListHandler := App.RegisterHandler(
			Services.BookListEndPoint(Services.NewBookService()),
			Services.CreateBookListRequest(),
			Services.CreateBookListResponse())
		v1.Handle(http.MethodGet, "books", bookListHandler)
	}

	router.Run(AppInit.Address)
}
