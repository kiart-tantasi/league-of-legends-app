# ############ CLIENT ############
FROM node:18.19.0-slim AS CLIENT_BUILD
WORKDIR /app
COPY ./client/package*.json /app/
RUN npm install
COPY ./client/. /app/
RUN ls
RUN npm run build

# ############ API ############
FROM openjdk:17.0.2-jdk-slim-buster
WORKDIR /app
COPY ./api/. /app/
COPY --from=CLIENT_BUILD /app/build/.  /app/src/main/resources/static/
RUN ./gradlew clean assemble
ARG SPRING_PROFILES_ACTIVE
ENV SPRING_PROFILES_ACTIVE=$SPRING_PROFILES_ACTIVE

ENTRYPOINT java -jar /app/build/libs/lol-api-0.0.1-SNAPSHOT.jar
