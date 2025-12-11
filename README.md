# social-ai

Full-stack social feed prototype with a React frontend and a Go backend.

## Project layout
- `src/`, `public/`, `package.json`: React SPA (Create React App).
- `backend/`: Go service (module `github.com/Pengc5Chen/social-ai/backend`).
  - `main.go`: entrypoint (HTTP on `:8080`).
  - `handler/`: HTTP handlers and routing (JWT auth).
  - `service/`: business logic.
  - `backend/`: Elasticsearch + Google Cloud Storage clients.
  - `constants/`: indices and connection constants.
  - `model/`: request/response models.
  - `app.yaml`: App Engine config.

## Requirements
- Node.js 18+ and npm (frontend).
- Go 1.22+ (backend).
- Elasticsearch reachable at `constants.ES_URL`.
- Google Cloud Storage bucket matching `constants.GCS_BUCKET`.

## Frontend (React)
Install dependencies:
```bash
npm install
```
Run dev server:
```bash
npm start
# http://localhost:3000
```
Build for production:
```bash
npm run build
```
Tests (if added):
```bash
npm test
```

## Backend (Go)
Install deps (module-aware):
```bash
cd backend
go mod download
```
Run locally:
```bash
go run .
# listens on :8080
```
Run tests (none yet):
```bash
go test ./...
```

### Config notes
- Connection info is currently hardcoded in `backend/constants/constants.go` (ES URL, credentials, GCS bucket, JWT signing key in handlers). For production, move these to env vars or secret manager before deploy.
- App Engine users: `backend/app.yaml` declares basic runtime settings; adjust service name, env vars, and scaling as needed.

## API sketch
- `POST /signup` — create user.
- `POST /signin` — returns JWT.
- `POST /upload` — authenticated, multipart file + `message`; stores media in GCS and metadata in ES.
- `GET /search?user=...` or `?keywords=...` — search posts.
- `DELETE /post/{id}` — authenticated delete.

## Deployment checklist
- Set real secrets via env (ES creds, JWT key, GCS bucket).
- Ensure ES indices exist (created automatically on start if missing).
- Verify service account has GCS write/read and ES access.
