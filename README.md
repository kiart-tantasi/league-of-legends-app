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

# Environment variables

## Client

### Development

```
# client/.env
REACT_APP_API_DOMAIN=http://localhost:8080 # where you run backend api at
DISABLE_ESLINT_PLUGIN=true # to disable lintint
```

# Todo

## Client

- [x] Set up tailwind

- [x] Set up eslint

- [ ] Set up github actions

- [ ] Configure `public/manifest.json` as below

  ```
  {
    "short_name": "React App",
    "name": "Create React App Sample",
    "icons": [
      {
        "src": "favicon.ico",
        "sizes": "64x64 32x32 24x24 16x16",
        "type": "image/x-icon"
      },
      {
        "src": "logo192.png",
        "type": "image/png",
        "sizes": "192x192"
      },
      {
        "src": "logo512.png",
        "type": "image/png",
        "sizes": "512x512"
      }c
    ],
    "start_url": ".",
    "display": "standalone",
    "theme_color": "#000000",
    "background_color": "#ffffff"
  }
  ```

## Backend

- [ ] Hot-reload

# Docker image

Build

```
docker build -t lol-app .
```

Run

```
docker run -dp 8080:8080 lol-app
```
