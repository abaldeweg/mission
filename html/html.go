package html

import (
	"baldeweg/mission/logfile"
	"baldeweg/mission/parseJson"
	"bytes"
	"html/template"
	"log"
	"time"
)

const tpl = `<ul>
    {{- range .Missions -}}
    <li>{{ formatDate .Date }} {{ getUnit .Unit }}: {{ .Situation }}, {{ .Location }}</li>
    {{- end -}}
</ul>`

func init() {
    log.SetPrefix("html: ")
    log.SetFlags(0)
}

func Export() string {
    var b bytes.Buffer
    logfile := parseJson.Read(string(logfile.ReadLogfile()))

    t, err := template.New("export").Funcs(template.FuncMap{
        "formatDate": formatDate,
        "getUnit": getUnit,
    }).Parse(tpl)
	if err != nil {
		log.Fatal(err)
	}

	if err = t.Execute(&b, logfile); err != nil {
        log.Fatal(err)
    }

    return b.String()
}

func formatDate(val string) string {
    t, err := time.Parse("2006-01-02", val)
    if err != nil {
        log.Fatal(err)
    }

    return t.Format("02.01.2006")
}

func getUnit(val string) string {
    missions := parseJson.Read(string(logfile.ReadLogfile()))

    return missions.Replacements[val]
}
