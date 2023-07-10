package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReadRoot(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
