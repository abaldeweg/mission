package create

import (
	"baldeweg/mission/logfile"
	"fmt"
	"log"
	"time"
)

func init() {
    log.SetPrefix("create: ")
    log.SetFlags(0)
}

func Create() {
    create := logfile.Mission{
        Date: time.Now().Format("2006-01-02"),
        Time: time.Now().Format("15:04"),
    }

    t := logfile.ParseYAML(logfile.ReadLogfile())
    t.Missions = append(t.Missions, create)

    logfile.WriteLogfile(logfile.WriteYAML(t))
    fmt.Printf("A new mission was created. Edit the details in the log file %s.\n", logfile.Path())
}
