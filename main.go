package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	registerCustomValidators()
	r.GET("/", handleIndex)
	r.POST("/mutant", handleMutant)
	r.GET("/stats", handleStats)
	_ = r.Run()
}

func handleIndex(c *gin.Context) {
	c.String(http.StatusOK, `Work`)
}