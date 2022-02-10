package main

import (
	"baldeweg/mission/create"
	"baldeweg/mission/logfile"
	"fmt"
	"os"
)

var data = `
missions:
`

func main() {
    if !logfile.HasLogfile() {
        logfile.WriteLogfile(logfile.WriteYAML(logfile.ParseYAML([]byte(data))))
    }

    action := "new"
    if len(os.Args) >= 2 {
        action = os.Args[1]
    }

    switch action {
        case "new":
            create.Create()
        case "help":
            fmt.Println("baldeweg/mission")
            fmt.Println("")
            fmt.Println("Commands")
            fmt.Println("mission new - Adds a new mission")
            fmt.Println("mission help - Shows the help")
        default:
            create.Create()
    }
}
