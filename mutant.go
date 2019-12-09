package main

import (
	"github.com/gin-gonic/gin"
	mutant "github.com/jhonynet/dna"
	"net/http"
)

type RequestData struct {
	Dna []string `json:"dna" binding:"required,squareMatrix,validDnaCharacters"`
}

func handleMutant(c *gin.Context) {
	var requestData  RequestData
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if mutant.IsMutant(requestData.Dna) {
		c.JSON(http.StatusOK, gin.H{})
	} else {
		c.JSON(http.StatusForbidden, gin.H{})
	}
}
