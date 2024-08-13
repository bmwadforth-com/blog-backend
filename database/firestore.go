package database

import (
	"cloud.google.com/go/firestore"
	"context"
	util "github.com/bmwadforth-com/armor-go/src/util"
)

var DbConnection *firestore.Client

func HealthCheck() error {
	_, err := DbConnection.Collection("healthz").Documents(context.Background()).GetAll()
	if err != nil {
		util.LogError("database healthcheck failed: %v", err)
		return err
	}

	return nil
}
