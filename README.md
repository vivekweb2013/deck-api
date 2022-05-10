# Deck API
This is a simple golang application that exposes deck APIs useful for playing card games.

## How to run
- Install go version `1.18.1` or above.
- Clone this repository
- Prepare config file using below command
```shell
cp app.yaml .app.yaml
```
You can edit the config values inside `.app.yaml` file. The application uses this file for retrieving configuration values.

- Start the server directly using below command
```shell
go run main.go
```

Or, If you want to build & run executable then below run commands
```shell
go build
./deck-api
```

This will start the http server.
