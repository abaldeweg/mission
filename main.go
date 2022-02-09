package main

import (
    "fmt"
    "os"
    "baldeweg/mission/create"
    "baldeweg/mission/logfile"
)

var data = `
missions:
`

func main() {
    if !logfile.IsFile(logfile.Path()) {
        logfile.WriteFile(logfile.Path(), logfile.WriteYAML(logfile.ParseYAML([]byte(data))))
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
