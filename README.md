# Demo

<img src="/demos/demo-1.jpg" alt="App demo" height="800px" />

# About project

- This is a simple app to see League of Legends match history
- Client is React and API is Spring Boot
- API is being replaced by Go (`./go-api`)

# Roadmap

- Migrate spring to golang
- Search suggestion
- Graph of damage done and recieved (mvp)
- Home page improvement

# Environment variables

## Client

`client/.env`

```
# optional
DISABLE_ESLINT_PLUGIN=true # disable lintint while developing
REACT_APP_IS_MOCK=true # replace fetching api with mocks
```

## API (Go)

`go-api/.env`

```
# optional
RIOT_API_REGION_ACCOUNT=<account-region>
RIOT_API_REGION_MATCH=<match-region>
RIOT_MATCH_AMOUNT=<match-amount>

# required
RIOT_API_KEY=<riot-api-key>
```

### Production

To use production env file, you need to provide:

- Export these env vars in your machine (with any method you prefer)
  - `ENV=production`
  - `PROJECT=ROOT=<project-location>/go-api`
- Put all other env vars in `<project-location>/go-api/.env.production`

## API (Spring)

`api/src/main/resources/application*.properties`

```
# required
riot.api.key=<api-key> # riot api key retrieved from https://developer.riotgames.com/
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

```
cd api
./gradlew bootRun
# you can also use your IDE to run app and that method is signicantly faster
```

### Test API manually

- Health endpoint
  ```
  curl "http://localhost:8080/api/health" -I
  ```
- Matches API endpoint (or with Makefile `make test`)
  ```
  curl "http://localhost:8080/api/v1/matches?gameName=เพชร&tagLine=ARAM" -I
  ```
