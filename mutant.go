package main

import (
	"github.com/gin-gonic/gin"
	mutant "github.com/jhonynet/dna"
	"net/http"
)

// structure for req/res
type RequestData struct {
	Dna []string `json:"dna" binding:"required,squareMatrix,validDnaCharacters"`
}

func handleMutant(c *gin.Context) {
	var requestData RequestData
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	isMutant := mutant.IsMutant(requestData.Dna)
	_, err := createDna(requestData.Dna, isMutant)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if isMutant {
		c.String(http.StatusOK, "")
	} else {
		c.String(http.StatusForbidden, "")
	}
}