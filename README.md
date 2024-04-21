![Golang Mascot](assets/go.webp "Golang Mascot")

# Player Service
A simple golang service for storing and retrieving scores for given players. Credits to [Lean Go with Tests](https://quii.gitbook.io/learn-go-with-tests).

## Build
You can build the service by running:
```bash
go build
```

## Running the service
You can either run the compiled executable or use go tools:
```bash
go run .
```

## Running tests
You can run all the tests with:
```bash
go test ./...
```

# Using the service
After running the service, it will listen on port 3000 and can handle 2 operations: retrieve wins and record win.

## Retrieving wins
The endpoint for retrieving wins for a given `player_name` is `GET /players/{player_name}`.
```bash
curl localhost:3000/players/Floyd
```

## Recording a win
In order to record a new win for a given `player_name`, you can `POST /players/{player_name}`.
```bash
curl -XPOST -v localhost:3000/players/Floyd
```
