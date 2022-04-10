package web

import (
	"baldeweg/mission/filetypes"
	"baldeweg/mission/logfile"
	"baldeweg/mission/parseJson"
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
)

func init() {
    log.SetPrefix("web: ")
    log.SetFlags(0)
}

func Web() {
    http.HandleFunc("/api/list", makeHandler(listHandler, "GET"))
    http.HandleFunc("/api/create", makeHandler(createHandler, "POST"))
    http.HandleFunc("/api/update", makeHandler(updateHandler, "PUT"))

    log.Fatal(http.ListenAndServe(":8080", nil))
}

func listHandler(w http.ResponseWriter, r *http.Request) {
    c := string(parseJson.Write(parseJson.Read(string(logfile.ReadLogfile()))))
    io.WriteString(w, c)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
    create := filetypes.Mission{
        Date: time.Now().Format("2006-01-02"),
        Time: time.Now().Format("15:04"),
    }

    t := parseJson.Read(string(logfile.ReadLogfile()))
    t.Missions = append(t.Missions, create)

    logfile.WriteLogfile(parseJson.Write(t))

    c := string(parseJson.Write(filetypes.Msg{Msg: "SUCCESS"}))
    io.WriteString(w, c)
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
    body, err := io.ReadAll(r.Body)
    if err != nil {
        log.Fatal(err)
    }

    logfile.WriteLogfile(body)

    c := string(parseJson.Write(filetypes.Msg{Msg: "SUCCESS"}))
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

        auth := strings.Split(r.Header.Get("Authorization"), " ")
        if len(auth) == 2 {
            if _, err := checkToken(auth[1]); err != nil {
                w.WriteHeader(401)
                return
            }
        }

        fn(w, r)
    }
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
