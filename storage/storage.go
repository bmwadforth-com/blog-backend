package storage

import (
	"blog-backend/util"
	"bytes"
	"cloud.google.com/go/storage"
	"context"
	"io"
	"time"
)

func createCloudStorageClient(ctx context.Context) (*storage.Client, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		util.SLogger.Errorf("failed to create cloud storage client: %v", err)
		return nil, err
	}

	projectId := util.Config.ProjectId
	bucketName := util.Config.CloudStorageBucket

	bucket := client.Bucket(bucketName)

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	if err := bucket.Create(ctx, projectId, nil); err != nil {
		util.SLogger.Errorf("failed to create cloud storage client: %v", err)
		return nil, err
	}

	// it is the responsibility of the calling function to ensure that the connection is closed
	// defer client.Close()
	return client, nil
}

func streamFileUpload(object string, content []byte) error {
	bucket := util.Config.CloudStorageBucket
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		util.SLogger.Errorf("failed to create cloud storage client: %v", err)
		return err
	}
	defer client.Close()

	buf := bytes.NewBuffer(content)
	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	wc := client.Bucket(bucket).Object(object).NewWriter(ctx)
	wc.ChunkSize = 0 // note retries are not supported for chunk size 0.

	if _, err := io.Copy(wc, buf); err != nil {
		util.SLogger.Errorf("failed to upload: %v", err)
		return err
	}

	if err := wc.Close(); err != nil {
		util.SLogger.Errorf("failed to upload: %v", err)
		return err
	}

	return nil
}

/*
func CloudStorageHealthCheck() error {
	ctx := context.Background()
	client, err := createCloudStorageClient(ctx)
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
*/
