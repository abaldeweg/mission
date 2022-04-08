package logfile

import (
	"baldeweg/mission/parseJson"
	"baldeweg/mission/storage/bucket"
	"baldeweg/mission/storage/file"
	"log"
	"os"
)

func init() {
    log.SetPrefix("logfile: ")
    log.SetFlags(0)
}

func HasLogfile() bool {
    storage := os.Getenv("STORAGE")

    if storage == "file" {
        return file.Exists()
    }
    if storage == "bucket" {
        return bucket.Exists(os.Getenv("BUCKET_NAME"), "missions.json")
    }

    return false
}

func WriteLogfile(content []byte) {
    storage := os.Getenv("STORAGE")

    if storage == "file" {
        file.Write(content)
    }
    if storage == "bucket" {
        bucket.Write(os.Getenv("BUCKET_NAME"), "missions.json", string(content))
    }
}

func ReadLogfile() []byte {
    var d []byte
    storage := os.Getenv("STORAGE")

    if storage == "file" {
        d = file.Read()
    }
    if storage == "bucket" {
        d = bucket.Read(os.Getenv("BUCKET_NAME"), "missions.json")
    }

    return d
}

func CreateTemplate() {
    WriteLogfile(parseJson.Template())
}
