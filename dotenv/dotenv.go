package dotenv

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {
    log.SetPrefix("dotenv: ")
    log.SetFlags(0)
}

func Load() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
}
