package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type app struct {
}

func (a *app) home(c *gin.Context) {
	log.Println("home handler")
	c.HTML(http.StatusOK, "home", gin.H{})
}
