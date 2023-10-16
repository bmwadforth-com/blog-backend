package database

import (
	"blog-backend/util"
	"cloud.google.com/go/firestore"
	"context"
)

var databaseConnection *firestore.Client

func createClient(ctx context.Context) (*firestore.Client, error) {
	client, err := firestore.NewClientWithDatabase(ctx, util.Config.ProjectId, util.Config.FireStoreDatabase)
	if err != nil {
		util.SLogger.Errorf("failed to create firestore client: %v", err)
		return nil, err
	}

	databaseConnection = client

	// it is the responsibility of the calling function to ensure that the connection is closed
	// defer client.Close()
	return client, nil
}

func HealthCheck() error {
	ctx := context.Background()
	client, err := createClient(ctx)
	defer client.Close()
	if err != nil {
		return err
	}

	_, err = client.Collection("healthz").Documents(ctx).GetAll()
	if err != nil {
		util.SLogger.Errorf("database healthcheck failed: %v", err)
		return err
	}

	return nil
}
