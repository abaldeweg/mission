package bucket

import (
	"context"
	"fmt"
	"io"
	"log"

	"cloud.google.com/go/storage"
)

func init() {
    log.SetPrefix("bucket: ")
    log.SetFlags(0)
}

func Write(bucket, object, content string) {
    ctx, storageClient := client()
    defer storageClient.Close()

    bkt := storageClient.Bucket(bucket)
    obj := bkt.Object(object)

    w := obj.NewWriter(ctx)
    if _, err := fmt.Fprintf(w, content); err != nil {
        log.Fatal(err)
    }
    if err := w.Close(); err != nil {
        log.Fatal(err)
    }
}

func Read(bucket, object string) []byte {
    ctx, storageClient := client()
    defer storageClient.Close()

    bkt := storageClient.Bucket(bucket)
    obj := bkt.Object(object)

    r, err := obj.NewReader(ctx)
    if err != nil {
        log.Fatal(err)
    }
    data, err := io.ReadAll(r);
    if err != nil {
        log.Fatal(err)
    }

    return data
}

func Exists(bucket, object string) bool {
    ctx, storageClient := client()
    defer storageClient.Close()

    _, err := storageClient.Bucket(bucket).Object(object).Attrs(ctx)
    if err == storage.ErrObjectNotExist {
        return false
    }
    if err != nil {
        return false
    }

    return true
}

func client() (context.Context, *storage.Client) {
    ctx := context.Background()
	storageClient, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

    return ctx, storageClient
}
