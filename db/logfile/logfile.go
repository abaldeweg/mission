package logfile

import (
	"baldeweg/mission/storage/bucket"
	"baldeweg/mission/storage/file"
	"baldeweg/mission/util"
	"fmt"
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
        return bucket.Exists(os.Getenv("BUCKET_NAME"), "missions.yaml")
    }

    return false
}

func WriteLogfile(content []byte) {
    storage := os.Getenv("STORAGE")

    if storage == "file" {
        file.Write(content)
    }
    if storage == "bucket" {
        //lint:ignore S1025 String() does not work as expected
        bucket.Write(os.Getenv("BUCKET_NAME"), "missions.yaml", fmt.Sprintf("%s", content))
    }
}

func ReadLogfile() []byte {
    var d []byte
    storage := os.Getenv("STORAGE")

    if storage == "file" {
        d = file.Read()
    }
    if storage == "bucket" {
        d = bucket.Read(os.Getenv("BUCKET_NAME"), "missions.yaml")
    }

    return d
}

func CreateTemplate() {
    WriteLogfile(util.Template())
}
