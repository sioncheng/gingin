package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type app struct {
}

func (a *app) home(c *gin.Context) {
	c.HTML(http.StatusOK, "home", nil)
}
