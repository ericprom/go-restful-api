package main

import (
	"adaptercrp-api/handlers"

	"github.com/labstack/echo/v4"
)

func handleRequest() {
	e := echo.New()
	e.GET("/", handlers.HomePage)
	e.GET("/api/articles", handlers.GetArticles)
	e.GET("/api/articles/:id", handlers.GetArticle)
	e.POST("/api/articles", handlers.CreateArticle)
	e.PUT("/api/articles/:id", handlers.UpdateArticle)
	e.DELETE("/api/articles/:id", handlers.DeleteArticle)
	e.Logger.Fatal(e.Start(":8081"))
}

func main() {
	handleRequest()
}
