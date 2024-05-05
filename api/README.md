# Spring project's README.md

**NOTE:** Spring project is now deprecated and is replaced by Go (./go-api), sadly

## Environment variables

`api/src/main/resources/application*.properties`

```
# required
riot.api.key=<api-key> # riot api key retrieved from https://developer.riotgames.com/
```

## Run app locally

```
./gradlew bootRun
# you can also use your IDE to run app and that method is signicantly faster
```

## Hot-reload on Intellij with `spring-boot-devtools`

You need to mark these settings

- Build, Execution, Deployment
  - Compiler
    - Build project automatically
- Advanced Settings
  - Allow auto-make to start even if developed application is currently running

## Prevent Builder and Constructor annotation

This project prevents using `@Builder` and `@*Constructor` in some places to prevent exposing internal representation so you will see a lot of manually written constructor
