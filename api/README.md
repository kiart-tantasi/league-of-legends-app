# Spring project's README.md

## Hot-reload on Intellij with `spring-boot-devtools`

You need to mark these settings

- Build, Execution, Deployment
  - Compiler
    - Build project automatically
- Advanced Settings
  - Allow auto-make to start even if developed application is currently running

## Prevent Builder and Constructor annotation

This project prevents using `@Builder` and `@*Constructor` in some places to prevent exposing internal representation so you will see a lot of manually written constructor
