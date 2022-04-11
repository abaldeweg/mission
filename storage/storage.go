package storage

import (
	"baldeweg/mission/storage/file"
	"baldeweg/mission/storage/gcpBucket"
	"log"
	"os"
)

type Adapter struct {
    Read func() []byte
    Write func([]byte)
    Exists func() bool
}

var Adapters = map[string]Adapter{
    "file":{file.Read, file.Write, file.Exists},
    "gcp-bucket":{gcpBucket.Read, gcpBucket.Write, gcpBucket.Exists},
}

func init() {
    log.SetPrefix("storage: ")
    log.SetFlags(0)
}

func Exists() bool {
    return existsHandler(Adapters[os.Getenv("STORAGE")].Exists)
}

func existsHandler(fn func() bool) bool {
    return fn()
}

func Write(content []byte) {
    writeHandler(Adapters[os.Getenv("STORAGE")].Write)(content)
}

func writeHandler(fn func([]byte)) func([]byte) {
    return func(content []byte) {
        fn(content)
    }
}

func Read() []byte {
    return readHandler(Adapters[os.Getenv("STORAGE")].Read)
}

func readHandler(fn func() []byte) []byte {
    return fn()
}
