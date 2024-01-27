# Demo

![App in mobile version](https://github.com/kiart-tantasi/league-of-legends-app/blob/main/demos/demo-1.png)

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

## API
```
# api/src/main/resources/application-local.properties
riot.api.key=<api-key>
```

# Run app locally

## Client

```
cd client
npm install # to install packages, run only at first time
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

# Todo

## Client

- [ ] Set up src path

## API

- [ ] Builder and check if it goes with Jackson
- [ ] Ccheckstyle
- [ ] GitHub Actions
- [ ] Swagger

# Production Image for ECS Deployment

Build
```
# prepare `./certificates` first
docker build -t league-of-legends-app .

# for macbook
docker build --platform=linux/amd64 -t league-of-legends-app .
```
