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

## How to run tests
- Run the following command from the root directory
```shell
go test -v -cover ./...
```

## API Endpoints
### `POST /api/v1/decks`
Create a new deck.
#### Params
| Name | Type | Optional | Description
| --- | --- | --- | --- |
| shuffled | boolean | yes | Can be set as `true/false`. If `true`, the cards from the deck will be shuffled.
| cards | csv | yes | CSV of card codes (e.g. `AS,KD,AC,2C,KH`). If provided, deck will be created with only specified cards otherwise deck will be created with all cards.

### `POST /api/v1/decks/:id/draw`
Draw card(s) from a specific deck.
#### Params
| Name | Type | Optional | Description
| --- | --- | --- | --- |
| id | uuid (string) | no | The uuid of deck.
| count | int | no | Number of cards to draw from the deck.
