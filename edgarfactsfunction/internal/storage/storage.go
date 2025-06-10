package storage

import (
	"bytes"
	"context"
	"io"
	"time"
	"errors"
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

var ErrNoFile error = errors.New("File does not exist")

func GetBytes(bucket, path string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	client, err := storage.NewClient(ctx)
		if err != nil {
			return []byte{}, err
		}
		defer client.Close()

		obj := client.Bucket(bucket).Object(path)

		_, err = obj.Attrs(ctx)
		if err != nil {
			if err == storage.ErrObjectNotExist {
				return []byte{}, ErrNoFile
			} else {
				return []byte{}, err
			}
		}

		reader, err := obj.NewReader(ctx)
		if err != nil {
			return []byte{}, err 
		}

		defer reader.Close()

		data, err := io.ReadAll(reader)
		if err != nil {
			return []byte{}, err
		}

		return data, nil
}