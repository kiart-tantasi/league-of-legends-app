# Demo

<img src="/demos/demo-1.png" alt="App demo" height="800px" />

# About project

- This is a simple app to see League of Legends match history
- Client is made with React
- API is made with Spring Boot

# Roadmap

- Search suggestion
- Graph of damage done and recieved
- Home page improvement

# Environment variables

## Client

`client/.env`

```
# optional
DISABLE_ESLINT_PLUGIN=true # disable lintint while developing
REACT_APP_IS_MOCK=true # replace fetching api with mocks
```

## API

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

Then `curl http://localhost:8080/api/health -I`

### Hot-reload on Intellij with `spring-boot-devtools`

You need to mark these settings

- Build, Execution, Deployment
  - Compiler
    - Build project automatically
- Advanced Settings
  - Allow auto-make to start even if developed application is currently running

# Prevent Builder and Constructor annotation

This project prevents using `@Builder` and `@*Constructor` in some places to prevent exposing internal representation so you will see a lot of manually written constructor
