# baldeweg/mission

A CLI to administer a log file of missions.

## Requirements

- [Go](https://go.dev/)
- Basic knowledge about the command line

## Getting Started

First, you need to install [Go](https://go.dev/).

Download the project archive from the [git repository](https://github.com/abaldeweg/mission).

Inside the project directory, you can build the app with `go build` command.

Run the command `mission`. Depending on the OS you need to add a file extension.

The app will create a log file and add a template for every new mission. The details needs to be added to the `missions.yaml` directly via an editor.

The `missions.yaml` file will be created in your home directory. Please, keep the indentation - that's very important for YAML to work correctly.

Find more about YAML:

- [https://yaml.org/](https://yaml.org/)
- [Wikipedia](https://en.m.wikipedia.org/wiki/YAML)

## Commands

- mission new - Adds a new mission
- mission help - Shows the help
