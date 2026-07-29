# League of Legends App

Match history viewer. Two independent apps live side-by-side:
- `client/` — React 18 + TypeScript frontend
- `goapi/` — Go (net/http) backend + mock Riot API

## Architecture

```
client/  →  GET /api/v1/matches?gameName=&tagLine=  →  goapi/  →  Riot API
```

Client proxies to `http://localhost:8080` in dev (set in `client/package.json`).

## Client (`client/`)

**Stack:** React 18, TypeScript, React Router v6, Tailwind CSS

**Key source layout:**
- `src/pages/SearchPage.tsx` — search form (route `*`)
- `src/pages/MatchPage.tsx` — match history display (route `/match`)
- `src/contexts/MatchContext.tsx` — global state provider
- `src/api/getMatchDetailList.ts` — API call to goapi
- `src/models/match.ts` — shared TypeScript types
- `src/utils/` — `search.ts`, `image.ts`, `populate.ts`
- `src/configs/ddragon.ts` — Data Dragon (LoL asset CDN) config
- `src/components/Layout.tsx`, `LoadingOverlay/` — shared UI

**Commands:**
```bash
cd client
npm install
npm start        # dev server on :3000
npm test         # run tests (no watch)
npm run build
npm run lint
npm run lint:fix
```

**Env:** copy `client/.env.example` → `client/.env`

## goapi (`goapi/`)

**Stack:** Go, standard `net/http`

**Two entrypoints:**
- `cmd/goapi/main.go` — main API server (default port 8080)
- `cmd/mockapi/main.go` — mock Riot API server (default port 8090)

**Internal packages:**
- `internal/match/` — handler, service, model, config, url builder
- `internal/health/` — health check handler
- `internal/middlewares/` — middleware chain: requestcontext → userdata → servertime → usercookie
- `internal/cache/` — optional MongoDB cache (toggled via `CACHE_ENABLED`)
- `internal/api/` — shared HTTP client
- `internal/mock/` — JSON fixtures used by mockapi
- `pkg/env/` — `.env` file loader

**Endpoints:**
- `GET /api/health`
- `GET /api/v1/matches?gameName=<name>&tagLine=<tag>`

**Commands:**
```bash
# Run locally (start mockapi first)
cd goapi && go run cmd/mockapi/main.go
cd goapi && go run cmd/goapi/main.go

# Build
cd goapi/cmd/goapi && go build
```

**Env:** copy `goapi/.env.example` → `goapi/.env`

For production, set `ENV=production` and `PROJECT_ROOT=<abs-path>/goapi` before running the binary (loads `goapi/.env.production`).

## Testing manually

```bash
curl "http://localhost:8080/api/health" -I
curl "http://localhost:8080/api/v1/matches?gameName=Hide%20on%20bush&tagLine=KR1"
```
