package main

import (
	"baldeweg/mission/dotenv"
	"baldeweg/mission/logfile"
	"baldeweg/mission/storage/file"
	"baldeweg/mission/web"
	"flag"
	"fmt"
	"log"
	"os"
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
    dotenv.Load()
    flag.Parse()

    file.SetPath(dir)
    if !logfile.HasLogfile() {
        logfile.CreateTemplate()
    }

    fmt.Println("baldeweg/mission https://github.com/abaldeweg/mission")

    web.Web()
}
