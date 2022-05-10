# Deck API
This is a simple golang application that exposes deck APIs useful for playing card games.

## How to run
- Install go version `1.18.1` or above
- Clone this repository
- Prepare config file

The application reads the configurations from `.app.yaml` file or the environment variables. The `app.yaml` is a sample config that can be used to create the actual config file. Execute the following command to create a app config file.
```shell
cp app.yaml .app.yaml
```
You can edit the config values inside `.app.yaml` file.

- Start the server directly using below command
```shell
go run main.go
```

Or, If you want to build & run executable then below run commands.
```shell
go build
./deck-api
```

This will start the http server on the host on port mentioned in the `.app.yaml` config file.

## How to run tests
- Run the following command from the root directory
```shell
go test -v -cover ./...
```
This will execute all the tests and also prints the code coverage percentage.

## API Endpoints
### `POST /api/v1/decks`
Create a new deck.
#### Params
| Name | Type | Optional | Description
| --- | --- | --- | --- |
| shuffled | boolean | yes | Can be set as `true/false`. If `true`, the cards from the deck will be shuffled.
| cards | csv | yes | CSV of card codes (e.g. `AS,KD,AC,2C,KH`). If provided, deck will be created with only specified cards otherwise deck will be created with all cards.

### `GET /api/v1/decks/:id`
Open a deck using deck uuid.
#### Params
| Param Name | Type | Optional | Description
| --- | --- | --- | --- |
| id | uuid (string) | no | The uuid of deck to be opened.

### `POST /api/v1/decks/:id/draw`
Draw card(s) from a specific deck.
#### Params
| Name | Type | Optional | Description
| --- | --- | --- | --- |
| id | uuid (string) | no | The uuid of deck.
| count | int | no | Number of cards to draw from the deck.
