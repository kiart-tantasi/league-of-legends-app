# Demo

<img src="/demos/demo-1.jpg" alt="App demo" height="800px" />

# About project

- This is a simple app to see League of Legends match history
- Client is written with React and API is written with Go (net/http)
- The API needs to connect to Riot API which you can apply for API key at https://developer.riotgames.com
  - This repo also provides mockapi for convenience. please see `Run app locally` section below
- The API was originally written with Spring (Java) but was replaced with Go (2024/05/05)
  - The deprecated Spring (Java) version can still be accessed [here](https://github.com/kiart-tantasi/deprecated-league-of-legends-app-api)

# Environment variables

## Client

Please create `client/.env` following `client/.env.example`

## API

Please create `goapi/.env` following `goapi/.env.example`

### Production env file (`goapi/.env.production`)

To use `goapi/.env.production`, you need to export 2 env vars below e.g. inline command

```
ENV=production
PROJECT_ROOT=<project-location>/goapi
```

Example (inline command)

```
ENV=production PROJECT_ROOT=/home/league-of-legends-app/goapi ./goapi
```

# Run app locally

_Please set up environment variables before proceeding_

## Client

```
cd client
npm install
npm start
```

## API

### First, Run mockapi (mock Riot API)

```
cd goapi
go run cmd/mockapi/main.go
```

### Second, Run goapi

```
cd goapi
go run cmd/goapi/main.go
```

### Test API manually

- Web app: http://localhost:3000

- Curl health endpoint

  ```
  curl "http://localhost:8080/api/health" -I
  ```

- Curl matches API endpoint

  ```
  curl "http://localhost:8080/api/v1/matches?gameName=NOPEEEE&tagLine=nopeeeee" -I
  ```

# Build

## Client

```
cd client
npm run build
```

## API

```
cd goapi/cmd/goapi
go build
```
