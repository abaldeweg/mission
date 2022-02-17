package main

import (
	"baldeweg/mission/commands/create"
	"baldeweg/mission/commands/html"
	"baldeweg/mission/commands/list"
	"baldeweg/mission/db/logfile"
	"flag"
	"fmt"
	"log"
	"os"
)

var data = `
replacements:
missions:
`

var dir string

func init() {
    path, err := os.Getwd()
    if err != nil {
        log.Fatal(err)
    }

    flag.StringVar(&dir, "path", path, "Specify the directory where the data should be stored.")
}

func main() {
    flag.Parse()

    logfile.SetPath(dir)

    if !logfile.HasLogfile() {
        logfile.WriteLogfile(logfile.WriteYAML(logfile.ParseYAML([]byte(data))))
    }

    action := "ls"
    if len(os.Args) >= 2 {
        action = os.Args[1]
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
