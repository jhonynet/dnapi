package main

import "github.com/gin-gonic/gin"

func handleStats(c *gin.Context) {
	c.JSON(200, gin.H{
		"count_mutant_dna": 12,
		"count_human_dna":  12,
		"ratio":            12.5,
	})
}