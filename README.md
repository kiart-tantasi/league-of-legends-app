# About project

- This is a simple app to see League of Legends match history
- Client-side is made with React
- Backend is made with Spring Boot

# Run app locally

## Client

```
cd client
npm start
```

Then visit http://localhost:3000

## Backend

```
cd backend
./gradlew bootRun
# you can also use your IDE to run app and that method is signicantly faster
```

Then `curl http://localhost:8080/api/health -I`

### Hot-reload with `spring-boot-devtools`

You need to use Intellij and set these 2 settings

- Build, Execution, Deployment
    - Compiler
        - Build project automatically
- Advanced Settings
    - Allow auto-make to start even if developed application is currently running

# Environment variables

## Client

### Development

```
# client/.env
REACT_APP_API_DOMAIN=http://localhost:8080 # where you run backend api at
DISABLE_ESLINT_PLUGIN=true # to disable lintint
```

## Backend
```
# backend/src/main/resources/application-local.properties
riot.api.key=<api-key>
```

# Todo

## Client

- [x] Set up tailwind
- [x] Set up eslint
- [ ] Set up github actions for jest and linting

## Backend

- [x] Hot-reload
- [x] Map response to java object
- [ ] controller advice
- [ ] Set up checkstyle
- [ ] Set up github actions for test and checkstyle

# Docker image

Build

```
docker build -t lol-app .
```

Run

```
docker run -dp 8080:8080 lol-app
```
