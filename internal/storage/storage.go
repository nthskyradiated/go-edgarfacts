package storage

import (
	"bytes"
	"context"
	"io"
	"time"

	"cloud.google.com/go/storage"
)

func UploadBytes(data []byte, bucket, path string) error {

	buffer := bytes.NewBuffer(data)

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, 120*time.Second)
	defer cancel()
	writer := client.Bucket(bucket).Object(path).NewWriter(ctx)


	_, err = io.Copy(writer, buffer)
	if err != nil {
		return err
	}

	err = writer.Close()
	if err != nil {
		return err
	}
	return nil
}