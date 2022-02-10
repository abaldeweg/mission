# baldeweg/mission

A CLI to administer a log file of missions.

## Getting Started

First, you need to install [Go](https://go.dev/) and then, inside the project directory, you can build the app with `go build` command.

Run the command `mission`.

The app will create a log file and add a template for every new mission. The details needs to be added to the `missions.yaml` directly via an editor.

The `missions.yaml` file will be created in your home directory. Please, keep the indentation - that's very important for YAML to work correctly.

## Commands

- mission new - Adds a new mission
- mission help - Shows the help

## Resources

- <https://yaml.org/>
- <https://en.m.wikipedia.org/wiki/YAML>
