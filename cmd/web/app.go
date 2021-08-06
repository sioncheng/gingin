package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sioncheng/gingin/cmd/service"
)

type app struct {
	SnippetService service.SnippetService
}

func (a *app) Home(c *gin.Context) {
	log.Println("home handler")
	c.HTML(http.StatusOK, "home", gin.H{})
}

func (a *app) ShowSnippet(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, "bad id")
		return
	}
	log.Println("show snippet", id)
	snippet := a.SnippetService.LoadSnippet(id)
	c.HTML(http.StatusOK, "show", gin.H{
		"Title":   snippet.Title,
		"Id":      snippet.Id,
		"Content": snippet.Content,
		"Created": snippet.Created,
		"Expires": snippet.Expires,
	})
}
