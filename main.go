package main

import (
	"baldeweg/mission/commands/create"
	"baldeweg/mission/commands/html"
	"baldeweg/mission/commands/list"
	"baldeweg/mission/dotenv"
	"baldeweg/mission/logfile"
	"baldeweg/mission/storage/file"
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

    action := "ls"
    if len(flag.Args()) >= 1 {
        action = flag.Args()[0]
    }

    switch action {
    case "ls":
        list.List()
    case "new":
        create.Create()
    case "html":
        html.Export()
    case "help":
        fmt.Println("baldeweg/mission")
        fmt.Println("A baldeweg OpenSource project")
        fmt.Println("https://github.com/abaldeweg/mission")
        fmt.Println("")
        fmt.Println("Commands")
        fmt.Println("mission ls - Lists all missions")
        fmt.Println("mission new - Adds a new mission")
        fmt.Println("mission html - Exports the missions to an HTML file")
        fmt.Println("mission help - Shows the help")
    default:
        list.List()
    }
}
