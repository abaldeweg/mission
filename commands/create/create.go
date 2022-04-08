package create

import (
	"baldeweg/mission/filetypes"
	"baldeweg/mission/logfile"
	"baldeweg/mission/parseJson"
	"baldeweg/mission/storage/file"
	"log"
	"time"

	"github.com/fatih/color"
)

func init() {
    log.SetPrefix("create: ")
    log.SetFlags(0)
}

func Create() {
    create := filetypes.Mission{
        Date: time.Now().Format("2006-01-02"),
        Time: time.Now().Format("15:04"),
    }

    t := parseJson.Read(string(logfile.ReadLogfile()))
    t.Missions = append(t.Missions, create)

    logfile.WriteLogfile(parseJson.Write(t))

    success := color.New(color.FgGreen)
    success.Println("A new mission was created.")
    success.Printf("File: %s\n", file.GetUrl())
}
