# baldeweg/mission

An app to administer a log file of missions.

## Requirements

- [Go](https://go.dev/)

## Getting Started

First, you need to install [Go](https://go.dev/).

Download the project archive from the [git repository](https://github.com/abaldeweg/mission).

Inside the project directory, you can build the app with the `go build` command. If you have [GoReleaser](https://goreleaser.com/) installed, instead run `goreleaser build --snapshot --rm-dist`.

Run the command `mission`. Depending on the OS you need to add a file extension.

The app will create a log file and add a template for every new mission, where you can edit the missions.

## Storage

Create a `.env` file to define some settings.

```env
// .env

STORAGE=gcp-bucket
FILE_PATH=.
GCP_BUCKET_NAME=name
GOOGLE_APPLICATION_CREDENTIALS=service-account-file.json
CORS_ALLOW_ORIGIN=http://localhost:8081
```

- STORAGE - Can be `file` or `gcp-bucket`
- FILE_PATH - Path where to store the files, only for file storage
- GCP_BUCKET_NAME - If `gcp-bucket` was chosen as storage, then define the bucket name.
- GOOGLE_APPLICATION_CREDENTIALS - Key file, for auth and buckets
- CORS_ALLOW_ORIGIN - Allowed origins
