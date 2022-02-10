package list

import (
	"baldeweg/mission/logfile"
	"log"
	"os"

	"github.com/olekukonko/tablewriter"
)

func init() {
    log.SetPrefix("list: ")
    log.SetFlags(0)
}

func List() {
    table := tablewriter.NewWriter(os.Stdout)
    table.SetHeader([]string{"Date", "Time", "Unit", "Situation", "Location"})

    data := logfile.ParseYAML(logfile.ReadLogfile())
    for _, v := range data.Missions {
        table.Append([]string{v.Date, v.Time, v.Unit, v.Situation, v.Location})
    }

    table.Render()
}
