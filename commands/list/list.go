package list

import (
	"baldeweg/mission/logfile"
	"baldeweg/mission/util"
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
    table.SetHeader([]string{"Date", "Time", "Unit", "Keyword", "Situation", "Location"})

    data := util.ParseYAML(logfile.ReadLogfile())
    for _, v := range data.Missions {
        table.Append([]string{v.Date, v.Time, getUnit(v.Unit), v.Keyword, v.Situation, v.Location})
    }

    table.Render()
}

func getUnit(val string) string {
    missions := util.ParseYAML(logfile.ReadLogfile())

    return missions.Replacements[val]
}
