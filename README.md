# About project

- This is simple to get league of legends data from riot api.
- Client-side is made with React (CRA)
- Backend is made with Spring Boot

# Run app locally

## Client

```
cd client
npm start
```

Visit http://localhost:3000

## Backend

```
cd backend
./gradlew bootRun
# you can also use your IDE to run app and that method is signicantly faster
```

Visit http://localhost:8080/api/health (You will get blank page with status 200)

# Todo

## Client

[ ] Set up tailwind

[X] Set up eslint

[ ] Set up github actions

[ ] Configure `public/manifest.json` as below

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

[ ] Hot-reload
