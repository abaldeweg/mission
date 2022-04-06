package create

import (
	"baldeweg/mission/db/logfile"
	"baldeweg/mission/storage/file"
	"baldeweg/mission/util"
	"log"
	"time"

	"github.com/fatih/color"
)

func init() {
    log.SetPrefix("create: ")
    log.SetFlags(0)
}

func Create() {
    create := util.Mission{
        Date: time.Now().Format("2006-01-02"),
        Time: time.Now().Format("15:04"),
    }

    t := util.ParseYAML(logfile.ReadLogfile())
    t.Missions = append(t.Missions, create)

    logfile.WriteLogfile(util.WriteYAML(t))

    success := color.New(color.FgGreen)
    success.Println("A new mission was created.")
    success.Printf("File: %s\n", file.GetUrl())
}
