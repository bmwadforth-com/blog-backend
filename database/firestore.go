package database

import (
	"blog-backend/util"
	"cloud.google.com/go/firestore"
	"context"
)

var DbConnection *firestore.Client

func HealthCheck() error {
	_, err := DbConnection.Collection("healthz").Documents(context.Background()).GetAll()
	if err != nil {
		util.SLogger.Errorf("database healthcheck failed: %v", err)
		return err
	}

	return nil
}
