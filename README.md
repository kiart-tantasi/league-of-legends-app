# About project

- This is a simple app to see League of Legends match history
- Client is made with React
- API is made with Spring Boot

# Environment variables

## Client

### Development

```
# client/.env
# please remove comments at the end before using
REACT_APP_API_DOMAIN=http://localhost:8080 # where you run api at
DISABLE_ESLINT_PLUGIN=true # to disable lintint while developing
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

- [ ]

## API

- [ ] Use Optional to handle null
- [ ] Set up checkstyle
- [ ] GitHub Actions for testing, checkstyling
- [ ] Use controller advice

# Docker image

Build
```
docker build -t app .
```

Build (production)
```
# needs to prepare ".env.production" for client and "application-production.properties" for api
docker build --build-arg SPRING_PROFILES_ACTIVE=production -t app .
```

Run
```
docker run -dp 8080:8080 app
```
