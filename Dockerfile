# ############ CLIENT ############
FROM node:18.19.0-slim AS CLIENT
WORKDIR /app
COPY ./client/package*.json /app/
RUN npm install
COPY ./client/. /app/
RUN ls
RUN npm run build

# ############ API ############
FROM openjdk:17.0.2-jdk-slim-buster AS API
WORKDIR /app
COPY ./api/. /app/
COPY --from=CLIENT /app/build/.  /app/src/main/resources/static/
RUN ./gradlew clean assemble
ARG SPRING_PROFILES_ACTIVE
ENV SPRING_PROFILES_ACTIVE=$SPRING_PROFILES_ACTIVE

# ############ NGINX ############
FROM nginx:alpine
RUN apk update && apk add openjdk17
WORKDIR /app
COPY --from=API /app/build/libs/lol-api-0.0.1-SNAPSHOT.jar .
COPY nginx.conf /etc/nginx/nginx.conf
COPY certificates/ssl_certificate.pem /etc/ssl/ssl_certificate.pem
COPY certificates/ssl_certificate_key.pem /etc/ssl/ssl_certificate_key.pem
ENTRYPOINT nohup java -jar lol-api-0.0.1-SNAPSHOT.jar & nginx -g 'daemon off;'
