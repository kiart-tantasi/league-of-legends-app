# Demo

<img src="/demos/demo-1.jpg" alt="App demo" height="800px" />

# About project

- This is a simple app to see League of Legends match history
- Client is written with React and API is written with Go (net/http)
- The API was originally written with Spring (Java) but was replaced with Go (2024/05/05)
  - The deprecated Spring (Java) version can still be accessed here: https://github.com/kiart-tantasi/deprecated-league-of-legends-app-api

# Roadmap

- User data stored in MongoDB (will change to sql db for compatibility with growthbook or other data-analysis tools)
- Deploy a self-managed or Purchase a managed sql database for keeping user data
- Search suggestion
- Graph of damage done and recieved (mvp)
- Migrate React to Nextjs and deployment from EC2 to Vercel + CSRF token

# Environment variables

## Client

`client/.env`

```
# optional
DISABLE_ESLINT_PLUGIN=true # disable lintint while developing
REACT_APP_IS_MOCK=true # replace fetching api with mocks
```

## API

- `goapi/.env`
- `goapi/.env.production`

```
# optional (actual api urls and api key are required for production)
RIOT_API_REGION_ACCOUNT=<account-region>
RIOT_API_REGION_MATCH=<match-region>
RIOT_MATCH_AMOUNT=<match-amount>
RIOT_ACCOUNT_API_URL=<retrived-from-https://developer.riotgames.com/>
RIOT_MATCH_IDS_API_URL=<retrived-from-https://developer.riotgames.com/>
RIOT_MATCH_DETAIL_API_URL=<retrived-from-https://developer.riotgames.com/>
RIOT_API_KEY=<retrived-from-https://developer.riotgames.com/>
CACHE_ENABLED=<true|false>
CACHE_MONGODB_URI=<uri>
CACHE_MONGODB_DATABASE_NAME=<database-name>
```

### Production env file (`.env.production`)

To use `.env.production`, you need to:

- Export these env vars with any method you prefer e.g. profile file, inline command
  - Env vars
    - `ENV=production`
    - `PROJECT_ROOT=<project-location>/goapi`
  - Example (inline command)
    ```
    cd goapi/cmd/goapi
    go build
    ENV=production PROJECT_ROOT=/home/league-of-legends-app/goapi ./goapi
    ```

# Run app locally

## Client

```
cd client
npm install
npm start
```

Then visit http://localhost:3000

## API

### First, Run mockapi (mocking Riot API)

```
cd goapi
go run cmd/mockapi/main.go
```

### Seond, Run goapi

```
cd goapi
go run cmd/goapi/main.go
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

# User data

- request_log
  - request
    - request_id (primary key)
    - timestamp
    - url
    - query_parameters
    - method

  - user
    - user_id
    - ip
    - headers

  - response
    - status_code
    - server_time
    - headers

- experiment_assignment
  - user_id
  - timestamp
  - experiment_id
  - variation_id
  - request_id
