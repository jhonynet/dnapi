package main

import (
	"cloud.google.com/go/datastore"
	"context"
)

var dsClient *datastore.Client = nil

func initDsClient() {
	if dsClient == nil {
		c := context.Background()
		i, err := datastore.NewClient(c, "")
		if err != nil {
			panic(err.Error())
		}
		dsClient = i
	}
}

func closeDsClient() {
	err := dsClient.Close()
	if err != nil {
		panic(err.Error())
	}
}