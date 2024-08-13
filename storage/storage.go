package storage

import (
	"blog-backend/util"
	"bytes"
	"cloud.google.com/go/storage"
	"context"
	armorUtil "github.com/bmwadforth-com/armor-go/src/util"
	"io"
	"time"
)

func streamFileUpload(object string, content []byte) error {
	bucket := util.Config.CloudStorageBucket
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		armorUtil.LogError("failed to create cloud storage client: %v", err)
		return err
	}
	defer client.Close()

	buf := bytes.NewBuffer(content)
	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	wc := client.Bucket(bucket).Object(object).NewWriter(ctx)
	wc.ChunkSize = 0 // note retries are not supported for chunk size 0.

	if _, err := io.Copy(wc, buf); err != nil {
		armorUtil.LogError("failed to upload: %v", err)
		return err
	}

	if err := wc.Close(); err != nil {
		armorUtil.LogError("failed to upload: %v", err)
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
		util.LogError("database healthcheck failed: %v", err)
		return err
	}

	return nil
}
*/
