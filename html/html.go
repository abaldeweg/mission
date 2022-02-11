package html

import (
	"baldeweg/mission/logfile"
	"fmt"
	"html/template"
	"log"
	"os"
	"time"
)

const tpl = `<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width,initial-scale=1.0">
    <title>{{ .Title }}</title>
</head>
<body>
    <h1>{{ .Title }}</h1>
    <ul>
        {{- range .Missions -}}
        <li>{{ formatDate .Date }} {{ .Time }} {{ .Unit }}: {{ .Situation }}, {{ .Location }}</li>
        {{- end -}}
    </ul>
</body>
</html>
`

type T struct {
    Title string
    Missions []logfile.Mission
}

func init() {
    log.SetPrefix("html: ")
    log.SetFlags(0)
}

func Export() {
    missions := logfile.ParseYAML(logfile.ReadLogfile())

	render(T{
		Title: "Missions Log",
		Missions: missions.Missions,
	})

    fmt.Printf("Wrote the HTML to %s\n", getUrl())
}

func render(data T) {
    t, err := template.New("export").Funcs(template.FuncMap{
        "formatDate": formatDate,
    }).Parse(tpl)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

    f, err := os.Create(getUrl())
    if err != nil {
        log.Fatal(err)
    }

	err = t.Execute(f, data)
    if err != nil {
        log.Fatal(err)
    }
}

func formatDate(val string) string {
    t, err := time.Parse("2006-01-02", val)
    if err != nil {
        log.Fatal(err)
    }

    return t.Format("02.01.2006")
}

func getUrl() string {
    home, err := os.UserHomeDir()
    if err != nil {
        log.Fatal(err)
    }

    return home + "/missions.html"
}
