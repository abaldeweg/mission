# baldeweg/mission

A CLI to administer a log file of missions.

![Screenshot](screenshot.jpg)

## Requirements

- [Go](https://go.dev/)
- Basic knowledge about the command line

## Getting Started

First, you need to install [Go](https://go.dev/).

Download the project archive from the [git repository](https://github.com/abaldeweg/mission).

Inside the project directory, you can build the app with the `go build` command. If you have [GoReleaser](https://goreleaser.com/) installed, instead run `goreleaser build --snapshot --rm-dist`.

Run the command `mission`. Depending on the OS you need to add a file extension.

The app will create a log file and add a template for every new mission. The details needs to be added to the `missions.json` directly via an editor.

The `missions.json` file will be created in your working directory. Please, keep the indentation - that's very important for JSON to work correctly.

Find more about JSON:

- [json.org](https://www.json.org/json-en.html)
- [Wikipedia](https://en.m.wikipedia.org/wiki/JSON)

## Storage

Create a `.env` file to define where to store the logfile.

```env
// .env

STORAGE=gcp-bucket
FILE_PATH=.
GCP_BUCKET_NAME=name
GOOGLE_APPLICATION_CREDENTIALS=service-account-file.json
CORS_ALLOW_ORIGIN=http://localhost:8081
```

- STORAGE - Can be file or gcp-bucket
- FILE_PATH - Path where to store the files, only for file storage
- GCP_BUCKET_NAME - If bucket was chosen as storage, then define the bucket name.
- GOOGLE_APPLICATION_CREDENTIALS - Key File, for auth and buckets
- CORS_ALLOW_ORIGIN - Allowed origins

## Flags

- `--path` - Specify the directory where the data should be stored.
