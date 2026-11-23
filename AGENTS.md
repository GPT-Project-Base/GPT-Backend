# Agent Notes

## Quick commands
- `make setup` — run `go mod tidy`
- `make build` — builds `./gpt-backend` from `main.go`
- `make test` — runs `go test -v`
- `make run` — runs `go run main.go`

## CI
- Workflow: `.github/workflows/go.yml`.
- Runs on push and PR to `main`.
- Uses Go `1.26.11`.
- Commands: `go build -v ./...` then `go test -v ./...`.
- Note: CI builds all packages recursively (`./...`), while `make build` produces the `./gpt-backend` binary from `main.go`.

## Project structure
- Single Go module (`github.com/gpt-backend`), single `main` package.
- Entrypoint: `main.go` → `main()` starts an HTTP server on `:8080`.
- Wiring: `CreateServer(ctx, dbfile)` → `CreateAPIMux(ctx, dbfile)` → `CreateAPIEndpoints(mux, db)` + GORM/sqlite DB.

## Database
- Uses GORM with `github.com/glebarez/sqlite` (CGO-free sqlite).
- Default DB file: `main.db` at repo root.
- Tests use `test.db` (gitignored; a copy may already exist in the working tree).

## Testing
- Tests build the mux via `CreateAPIMux(context.Background(), "test.db")` and use `httptest.NewRecorder`; no real server is started.

## Notes
- `Dockerfile` and `TODO.md` are currently empty; do not rely on them.
- No lint config or README exists yet.
