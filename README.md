# Demo

<img src="/demos/demo-1.png" alt="App demo" height="800px" />

# About project

- This is a simple app to see League of Legends match history
- Client is made with React
- API is made with Spring Boot

# Environment variables

## Client

### Development

```
# client/.env
# below is config to disable lintint while developing
DISABLE_ESLINT_PLUGIN=true
# prevent fetching riot api with mocks
REACT_APP_IS_MOCK=true
```

## Production

```
# still no env for production client
```

## API
```
# api/src/main/resources/application-local.properties
# required
riot.api.key=<api-key>

# optional
riot.api.key.secondary=<api-key>
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

Then `curl http://localhost:8080/api/health -I`

### Hot-reload `spring-boot-devtools` on Intellij

You need to mark these settings

- Build, Execution, Deployment
    - Compiler
        - Build project automatically
- Advanced Settings
    - Allow auto-make to start even if developed application is currently running

# Production Image for ECS Deployment

Build
```
# prepare `./certificates` first
docker build -t league-of-legends-app .

# for macbook
docker build --platform=linux/amd64 -t league-of-legends-app .
```
