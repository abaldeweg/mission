package main

import (
	"baldeweg/mission/logfile"
	"baldeweg/mission/storage/file"
	"baldeweg/mission/web"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var dir string

func init() {
    path, err := os.Getwd()
    if err != nil {
        log.Fatal(err)
    }

    flag.StringVar(&dir, "path", path, "Specify the directory where the data should be stored.")
}

func main() {
    if _, err := os.Stat("./.env"); err == nil {
        if err := godotenv.Load(); err != nil {
            log.Fatal("Error loading .env file")
        }
    }

    flag.Parse()

    file.SetPath(dir)
    if !logfile.HasLogfile() {
        logfile.CreateTemplate()
    }

    fmt.Println("baldeweg/mission https://github.com/abaldeweg/mission")

    web.Web()
}
