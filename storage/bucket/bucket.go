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

func Write(content []byte) {
    ctx, storageClient := client()
    defer storageClient.Close()

    bkt := storageClient.Bucket(os.Getenv("BUCKET_NAME"))
    obj := bkt.Object("missions.json")

    w := obj.NewWriter(ctx)
    if _, err := fmt.Fprintf(w, string(content)); err != nil {
        log.Fatal(err)
    }
    if err := w.Close(); err != nil {
        log.Fatal(err)
    }
}

func Read() []byte {
    ctx, storageClient := client()
    defer storageClient.Close()

    bkt := storageClient.Bucket(os.Getenv("BUCKET_NAME"))
    obj := bkt.Object("missions.json")

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

func Exists() bool {
    ctx, storageClient := client()
    defer storageClient.Close()

    _, err := storageClient.Bucket(os.Getenv("BUCKET_NAME")).Object("missions.json").Attrs(ctx)
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
