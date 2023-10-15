package database

import (
	"cloud.google.com/go/firestore"
	"context"
	"blog-backend/util"
)

var databaseConnection *firestore.Client

type EDatabaseResult uint

var (
	DbresultNotFound   EDatabaseResult = 0
	DbresultError      EDatabaseResult = 1
	DbresultOk         EDatabaseResult = 2
	DbresultIncomplete EDatabaseResult = 3
)

func createClient(ctx context.Context) *firestore.Client {
	client, err := firestore.NewClientWithDatabase(ctx, util.Config.ProjectId, util.Config.FireStoreDatabase)
	if err != nil {
		util.SLogger.Errorf("failed to create firestore client: %v", err)
	}

	databaseConnection = client

	// it is the responsibility of the calling function to ensure that the connection is closed
	// defer client.Close()

	return client
}
