package parseJson

import (
	"encoding/json"
	"log"
)

type Logfile struct {
    Notes []string `json:"notes"`
    Replacements map[string]string `json:"replacements"`
    Missions []Mission `json:"missions"`
}

type Mission struct {
    Date string `json:"date"`
    Time string `json:"time"`
    Keyword string `json:"keyword"`
    Situation string `json:"situation"`
    Unit string `json:"unit"`
    Location string `json:"location"`
    Links []string `json:"links"`
}

var data = `{"notes":[],"replacements":{},"missions":[]}`

func init() {
    log.SetPrefix("parseJson: ")
    log.SetFlags(0)
}

func Read(blob string) Logfile {
    var d Logfile
	if err := json.Unmarshal([]byte(blob), &d); err != nil {
		log.Fatal(err)
	}

    return d
}

func Write(data interface{}) []byte {
	d, err := json.Marshal(&data)
    if err != nil {
        log.Fatal(err)
    }

    return d
}

func Template() []byte {
    return Write(data)
}
