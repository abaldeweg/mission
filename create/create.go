package create

import (
    "fmt"
    "log"
    "baldeweg/mission/logfile"
)

func init() {
    log.SetPrefix("create: ")
    log.SetFlags(0)
}

func Create() {
    logfile.WriteFile(logfile.Path(), addContent(logfile.ReadFile(logfile.Path())))
    fmt.Printf("A new mission was created. Edit the details in the log file %s.\n", logfile.Path())
}

func addContent(file []byte) []byte {
    t := logfile.ParseYAML(file)

    create := logfile.Mission{Date: "123", Links: []string{"url"}}

    t.Missions = append(t.Missions, create)

    return logfile.WriteYAML(t)
}
