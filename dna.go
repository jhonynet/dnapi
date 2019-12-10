package main

import (
	"cloud.google.com/go/datastore"
	"context"
	mutant "github.com/jhonynet/dna"
)

// structure for dna record
type DnaRecord struct {
	Id       string     `json:"id" datastore:"-"`
	Dna      mutant.Dna `json:"name"`
	IsMutant bool       `json:"isMutant"`
}

func getDnaKey(id string) *datastore.Key {
	return datastore.NameKey("Dna", id, nil)
}

// get dna by uid
func getDnaById(id string) (*DnaRecord, error) {
	var dna DnaRecord
	ctx := context.Background()
	k := getDnaKey(id)
	if err := dsClient.Get(ctx, k, &dna); err != nil {
		return nil, err
	}

	return &dna, nil
}

// create dna if not exists
func createDna(dna mutant.Dna, isMutant bool) (*DnaRecord, error) {
	ctx := context.Background()
	uid := mutant.BuildUniqueId(dna)
	record, err := getDnaById(uid)
	if err != nil && err.Error() != "datastore: no such entity" {
		return record, err
	}
	// if record exists avoid creation
	if record != nil && record.Id != "" {
		return record, nil
	}
	// create new record and save
	var newRecord = DnaRecord{
		Id:       uid,
		Dna:      dna,
		IsMutant: isMutant,
	}
	// save
	if _, err := dsClient.Put(ctx, getDnaKey(uid), &newRecord); err != nil {
		return &newRecord, err
	}

	return &newRecord, nil
}