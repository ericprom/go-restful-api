package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type (
	article struct {
		ID      int    `json:"ID"`
		Title   string `json:"Title"`
		Desc    string `json:"Desc"`
		Content string `json:"Content"`
	}
)

var (
	articles = map[int]*article{}
	seq      = 1
)

func getArticles(c echo.Context) error {
	return c.JSON(http.StatusOK, articles)
}

func getArticle(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, articles[id])
}

func createArticle(c echo.Context) error {
	a := &article{
		ID: seq,
	}
	if err := c.Bind(a); err != nil {
		return err
	}
	articles[a.ID] = a
	seq++
	return c.JSON(http.StatusCreated, a)
}

func updateArticle(c echo.Context) error {
	a := new(article)
	if err := c.Bind(a); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	articles[id].Title = a.Title
	articles[id].Desc = a.Desc
	articles[id].Content = a.Content
	return c.JSON(http.StatusOK, articles[id])
}

func deleteArticle(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(articles, id)
	return c.NoContent(http.StatusNoContent)
}

func homePage(c echo.Context) error {
	return c.String(http.StatusOK, "Homepage Endpoint Hit")
}

func handleRequest() {
	e := echo.New()
	e.GET("/", homePage)
	e.GET("/api/articles", getArticles)
	e.GET("/api/articles/:id", getArticle)
	e.POST("/api/articles", createArticle)
	e.PUT("/api/articles/:id", updateArticle)
	e.DELETE("/api/articles/:id", deleteArticle)
	e.Logger.Fatal(e.Start(":8081"))
}

func main() {
	handleRequest()
}
