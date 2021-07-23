package handlers

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

func GetArticles(c echo.Context) error {
	return c.JSON(http.StatusOK, articles)
}

func GetArticle(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, articles[id])
}

func CreateArticle(c echo.Context) error {
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

func UpdateArticle(c echo.Context) error {
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

func DeleteArticle(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(articles, id)
	return c.NoContent(http.StatusNoContent)
}
