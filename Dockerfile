# ############ CLIENT ############
FROM node:18.19.0-slim AS CLIENT_BUILD
WORKDIR /app
COPY ./client/package*.json /app/
RUN npm install
COPY ./client/. /app/
RUN ls
RUN npm run build

# ############ BACKEND ############
FROM openjdk:17.0.2-jdk-slim-buster
WORKDIR /app
COPY ./backend/. /app/
COPY --from=CLIENT_BUILD /app/build/.  /app/src/main/resources/static/
RUN ./gradlew clean assemble
ENTRYPOINT java -jar /app/build/libs/lol-api-0.0.1-SNAPSHOT.jar
