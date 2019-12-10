package main

import (
	"cloud.google.com/go/datastore"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

// structure for dna record
type StatsRecord struct {
	Id             string `json:"id"`
	CountMutantDna int    `json:"count_mutant_dna"`
	CountHumanDna  int    `json:"count_human_dna"`
}

// key for datastore
var statsKey = datastore.Key{
	Kind: "Stats",
	Name: "stats",
}

// /stats handler
func handleStats(c *gin.Context) {
	stats, err := getStats()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"count_mutant_dna": stats.CountMutantDna,
		"count_human_dna":  stats.CountHumanDna,
		"ratio":            calcRatio(stats),
	})
}

// safe calculation of ratio
func calcRatio(stats *StatsRecord) float64 {
	if stats.CountHumanDna > 0 {
		return float64(stats.CountMutantDna) / float64(stats.CountHumanDna)
	}

	return 0
}

// get stats from database
func getStats() (*StatsRecord, error) {
	var stats StatsRecord
	ctx := context.Background()
	err := dsClient.Get(ctx, &statsKey, &stats)
	if err != nil && err.Error() != "datastore: no such entity" {
		return &stats, err
	}
	if err != nil && err.Error() == "datastore: no such entity" {
		stats.CountHumanDna = 0
		stats.CountMutantDna = 0
		_, err := dsClient.Put(ctx, &statsKey, stats)
		if err != nil {
			return &stats, err
		}
	}

	return &stats, nil
}

// increase stats by 1
func IncreaseStat(kind string) error {
	stats, err := getStats()
	if err != nil {
		return err
	}
	if kind == "human" {
		stats.CountHumanDna++
	} else {
		stats.CountMutantDna++
	}
	ctx := context.Background()
	_, err = dsClient.Put(ctx, &statsKey, stats)
	if err != nil {
		return err
	}

	return nil
}