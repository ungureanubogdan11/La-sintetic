# 2026-11-26

- created frontend in `src-client/` (React 19 + TypeScript + Tailwind v4, scaffolded with Vite 8)

- setup backend structure in `src/` (Go 1.25 + Gin + GORM + PostgreSQL + JWT + bcrypt)
  - `src/cmd/api/main.go` 
    - entry point: loads config, connects to the DB, wires the handlers, starts Gin, shuts down gracefully
  - `src/internal/config/config.go` 
    - reads and validates env vars (PG_*, PORT, ENV, JWT_*) and builds the DB connection string
  - `src/internal/database/postgres.go`
    - opens the GORM connection to PostgreSQL, sets pool limits and pings at startup
  - `src/internal/apperror/apperror.go`
    - HTTP-agnostic domain errors (invalid, unauthorized, forbidden, not found, conflict, internal) returned by services
  - `src/internal/httpx/response.go` 
    - standard JSON envelope and mapping of apperror kinds to HTTP status codes
    - success: `{"data": ..., "meta": {...}}` (meta only for paginated lists)
    - error: `{"error": {"code": "...", "message": "...", "details": ...}}`
  - `src/internal/middleware/recovery.go`
    - turns panics into a standard JSON 500 response
  - `src/internal/handlers/`
    - `GET /api/health`: reports API status and database reachability
  - `src/internal/models/user.go`
    - GORM `User` entity plus the Role and Position enums, mapped to the `users` table
  - `src/internal/services/` 
    - business-logic layer (sign-ups, waitlist, cost per player)
  - `src/internal/repositories/`
    - data-access layer (all GORM queries go here)- `src-client/` (React 19 + TypeScript + Tailwind v4, scaffolded with Vite 8)