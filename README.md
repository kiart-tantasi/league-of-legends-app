# Demo

<img src="/demos/demo-1.jpg" alt="App demo" height="800px" />

# About project

- This is a simple app to see League of Legends match history
- Client is React and API is Go
- Spring was written with Spring and got fully replaced by Go (2024/05/05)

# Roadmap

- Compare net/http and gin
- Write tests for Go (either net/http or gin)
- Search suggestion
- Migrate React to Nextjs and deployment from EC2 to Vercel
- Graph of damage done and recieved (mvp)

# Environment variables

## Client

`client/.env`

```
# optional
DISABLE_ESLINT_PLUGIN=true # disable lintint while developing
REACT_APP_IS_MOCK=true # replace fetching api with mocks
```

## API (Go)

- `go-api/.env`
- `go-api/.env.production`

```
# optional (actual api urls and api key are required for production)
RIOT_API_REGION_ACCOUNT=<account-region>
RIOT_API_REGION_MATCH=<match-region>
RIOT_MATCH_AMOUNT=<match-amount>
RIOT_ACCOUNT_API_URL=<retrived-from-https://developer.riotgames.com/>
RIOT_MATCH_IDS_API_URL=<retrived-from-https://developer.riotgames.com/>
RIOT_MATCH_DETAIL_API_URL=<retrived-from-https://developer.riotgames.com/>
RIOT_API_KEY=<retrived-from-https://developer.riotgames.com/>
```

### Production env file (`.env.production`)

To use `.env.production`, you need to:

- Export these env vars with any method you prefer e.g. profile file, inline command
  - env vars
    - `ENV=production`
    - `PROJECT_ROOT=<project-location>/go-api`
  - example (inline command)
    ```
    go build
    ENV=production PROJECT_ROOT=/league-of-legends-app/go-api ./go-api
    ```

# Run app locally

## Client

```
cd client
npm install
npm start
```

Then visit http://localhost:3000

## API (Go)

### First, Run mock-api (mocking Riot API)

```
cd go-api
go run cmd/mock-api/main.go
```

### Seond, Run go-api

```
cd go-api
go run cmd/go-api/main.go
```

### Test API manually

- Health endpoint
  ```
  curl "http://localhost:8080/api/health" -I
  ```
- Matches API endpoint (or with Makefile `make test`)
  ```
  curl "http://localhost:8080/api/v1/matches?gameName=%E0%B9%80%E0%B8%9E%E0%B8%8A%E0%B8%A3&tagLine=ARAM" -I
  ```
