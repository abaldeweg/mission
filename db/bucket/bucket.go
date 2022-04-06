package bucket

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

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

func Read(bucket, object string) {
    ctx, storageClient := client()
    defer storageClient.Close()

    bkt := storageClient.Bucket(bucket)
    obj := bkt.Object(object)

    r, err := obj.NewReader(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if _, err := io.Copy(os.Stdout, r); err != nil {
        log.Fatal(err)
    }
}

func client() (context.Context, *storage.Client) {
    ctx := context.Background()
	storageClient, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

    return ctx, storageClient
}
