package parseJson

import (
	"baldeweg/mission/filetypes"
	"encoding/json"
	"log"
)

var data = `{"notes":[],"replacements":{},"missions":[]}`

func init() {
    log.SetPrefix("parseJson: ")
    log.SetFlags(0)
}

func Read(blob string) filetypes.Logfile {
    var d filetypes.Logfile
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
