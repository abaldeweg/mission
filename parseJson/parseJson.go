package parseJson

import (
	"encoding/json"
	"log"
)

type Logfile struct {
    Notes []string
    Replacements map[string]string
    Missions []Mission
}

type Mission struct {
    Date string
    Time string
    Keyword string
    Situation string
    Unit string
    Location string
    Links []string
}

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

func Write(data Logfile) []byte {
	d, err := json.Marshal(&data)
    if err != nil {
        log.Fatal(err)
    }

    return d
}
