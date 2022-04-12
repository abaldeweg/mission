package web

import (
	"baldeweg/mission/export/html"
	"baldeweg/mission/filetype"
	"baldeweg/mission/storage"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
)

var filename = "missions.json"

func init() {
    log.SetPrefix("web: ")
    log.SetFlags(0)
}

type Msg struct {
    Msg string `json:"msg"`
}

func Web() {
    http.HandleFunc("/api/show", makeHandler(showHandler, "GET"))
    http.HandleFunc("/api/create", makeHandler(createHandler, "POST"))
    http.HandleFunc("/api/update", makeHandler(updateHandler, "PUT"))
    http.HandleFunc("/api/export/html", makeHandler(htmlExportHandler, "GET"))

    log.Fatal(http.ListenAndServe(":8080", nil))
}

func showHandler(w http.ResponseWriter, r *http.Request) {
    c := string(storage.Read(filename))
    io.WriteString(w, c)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
    create := filetype.Mission{
        Date: time.Now().Format("2006-01-02"),
        Time: time.Now().Format("15:04"),
    }

    t := unmarshalJson(string(storage.Read(filename)))
    t.Missions = append(t.Missions, create)

    storage.Write(filename, marshalJson(t))

    c := string(marshalJson(Msg{Msg: "SUCCESS"}))
    io.WriteString(w, c)
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
    body, err := io.ReadAll(r.Body)
    if err != nil {
        log.Fatal(err)
    }

    storage.Write(filename, body)

    c := string(marshalJson(Msg{Msg: "SUCCESS"}))
    io.WriteString(w, c)
}

func htmlExportHandler(w http.ResponseWriter, r *http.Request) {
    c := string(marshalJson(filetype.Export{Type: "html", Body: html.Export()}))
    io.WriteString(w, c)
}

func makeHandler(fn func(http.ResponseWriter, *http.Request), method string) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != method && r.Method != "OPTIONS" {
            http.NotFound(w, r)
            return
        }

        w.Header().Set("Access-Control-Allow-Origin", os.Getenv("CORS_ALLOW_ORIGIN"))
        w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
        w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
        w.Header().Set("Content-Type", "application/json")

        if r.Method == "OPTIONS" {
            return
        }

        if !isAuthenticated(r.Header.Get("Authorization")) {
            w.WriteHeader(401)
            return
        }

        fn(w, r)
    }
}

func isAuthenticated(auth string) bool {
    token := strings.Split(auth, " ")
    if len(token) == 2 {
        if _, err := checkToken(token[1]); err == nil {
            return true
        }
    }

    return false
}

func checkToken(idToken string) (*auth.Token, error) {
    ctx := context.Background()

    app, err := firebase.NewApp(ctx, nil)
    if err != nil {
        return nil, err
    }

    client, err := app.Auth(ctx)
    if err != nil {
        return nil, err
    }

    token, err := client.VerifyIDTokenAndCheckRevoked(ctx, idToken)
    if err != nil {
        return nil, err
    }

    return token, nil
}

func unmarshalJson(blob string) filetype.Logfile {
    var d filetype.Logfile
	if err := json.Unmarshal([]byte(blob), &d); err != nil {
		log.Fatal(err)
	}

    return d
}

func marshalJson(data interface{}) []byte {
	d, err := json.Marshal(&data)
    if err != nil {
        log.Fatal(err)
    }

    return d
}
