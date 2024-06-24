# NOTE: app is not currently deployed in container so this file is not valid

# ############ API ############
# TODO: build go here

# ############ CLIENT ############
FROM node:18.19.0-slim AS CLIENT
WORKDIR /app
COPY ./client/package*.json /app/
RUN npm install
COPY ./client/. /app/
RUN npm run build

# ############ NGINX ############
FROM nginx:alpine
# TODO: install go here
WORKDIR /app
# TODO: copy api build to here
COPY nginx.conf /etc/nginx/nginx.conf
COPY certificates/ssl_certificate.pem /etc/ssl/ssl_certificate.pem
COPY certificates/ssl_certificate_key.pem /etc/ssl/ssl_certificate_key.pem
COPY --from=CLIENT /app/build /usr/share/nginx/html/
ENTRYPOINT echo 'TODO: put command to start api here' & nginx -g 'daemon off;'
