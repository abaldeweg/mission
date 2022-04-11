package html

import (
	"baldeweg/mission/filetype"
	"baldeweg/mission/storage"
	"bytes"
	"encoding/json"
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
    storage := Read(string(storage.Read()))

    t, err := template.New("export").Funcs(template.FuncMap{
        "formatDate": formatDate,
        "getUnit": getUnit,
    }).Parse(tpl)
	if err != nil {
		log.Fatal(err)
	}

	if err = t.Execute(&b, storage); err != nil {
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
    missions := Read(string(storage.Read()))

    return missions.Replacements[val]
}

func Read(blob string) filetype.Logfile {
    var d filetype.Logfile
	if err := json.Unmarshal([]byte(blob), &d); err != nil {
		log.Fatal(err)
	}

    return d
}
