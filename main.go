package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"web-scene/App"
	"web-scene/App/Services"
	"web-scene/AppInit"
)

func main() {
	simpleLog()

	router := gin.Default()

	v1 := router.Group("v1")
	{
		bookListHandler := App.RegisterHandler(
			Services.BookListEndPoint(Services.NewBookService()),
			Services.CreateBookListRequest(),
			Services.CreateBookListResponse())
		v1.Handle(http.MethodGet, "books", bookListHandler)

		bookDetailHandler := App.RegisterHandler(
			Services.BookDetailEndPoint(Services.NewBookService()),
			Services.CreateBookRequest(),
			Services.CreateBookResponse(),
		)
		v1.Handle(http.MethodGet, "books/:id", bookDetailHandler)
	}

	router.Run(AppInit.Address)
}

func simpleLog() {
	logFile, err := os.OpenFile("gin-log.log", os.O_CREATE|os.O_APPEND, 666)
	gin.DefaultWriter = io.MultiWriter(logFile)
	if err != nil {
		log.Fatalln("created log file failed", err.Error())
	}
}
