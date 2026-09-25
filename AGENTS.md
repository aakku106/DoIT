# DoIT

Go CLI todo app backed by SQLite through `ncruces/go-sqlite3` (WASM-based, no CGO). The only configured application is `cmd/todo`; dev and release builds do not include `temp/`.

## Toolchain and Verification

- `go.mod` requires Go 1.25.6 or newer.
- Focused application build: `go build -o doit ./cmd/todo`.
- Run without installing: `go run ./cmd/todo <command>`; every non-`init` command still needs an initialized project.
- `./dev-build.sh` must run from the repository root and writes `~/go/bin/xdoit`. It assumes that directory and `PATH` are already set up, and its final command masks `go build` failures, so do not trust its exit status alone.
- The README's `go install github.com/aakku106/DoIT@latest` does not match this layout: the module root has no Go package.
- There are no tests, lint/format/typecheck configurations, task runner, pre-commit hooks, or PR CI. There is consequently no single-test command; use the focused build above after changes.

## SQL and Code Generation

- `sql/schema.sql` and `sql/queries.sql` are the sources of truth. After changing either, run `sqlc generate` from the repository root.
- `sqlc generate` writes `internal/store/`; never hand-edit those generated files. The `sqlc` executable is not pinned, while the current output identifies sqlc v1.31.1; use that version to avoid unrelated churn.
- The schema is separately embedded through the intentionally existing but misspelled `sql/assects.go`. Only `doit init` executes it; normal database opening does not apply or upgrade a schema.
- `migrations/001_init.sql` is stale and unreferenced. Existing databases have no migration path in the current CLI.

## Runtime Boundaries

- `cmd/todo/main.go` routes arguments; most command handlers live beside it in package `main`. `internal/cli` performs operations with a concrete `*store.Queries`, prints directly, and often exits instead of returning errors, so do not treat it as a reusable library API.
- `internal/db.NewSQLite` walks from the working directory to the nearest ancestor containing `.doit/`, then opens `.doit/doit.db`. `doit init` always creates `.doit/` in the current directory and refuses if that exact directory already exists.
- This repository already contains a tracked `.doit/`, so run CLI smoke tests from a clean temporary project directory rather than accidentally modifying the checked-in database.
- `temp/` is an experimental executable, and `internal/todo/service.go` is a dead stub; neither is part of the configured application build.

## Data-Layer Semantics and Known Hazards

- List labels are 0-based display positions, not database IDs. Mutation handlers re-query and index ordered IDs at execution time; `add` misleadingly prints the newly assigned database ID. Ordering has no secondary tie-breaker for equal timestamps.
- Every live path hardcodes session `"todo"`. `sessionCall` in `cmd/todo/session.go` is empty, and clear queries are unscoped `DELETE FROM <table>` statements, so any future session work must also scope destructive queries.
- State changes are separate copy and delete operations, not database transactions; generated names such as `CompleteTodoTransaction` do not start a transaction.
- Position bounds checks are currently off by one. Invalid `done`/move input can panic or affect item 0, and `trash remove`, `trash nuke`, and `ignored nuke` currently call completed-table handlers.
- Parsed `-t` deadlines are discarded on insertion and expiration is not enforced.

## Release

- The only workflow runs on pushed tags matching `v*`; it publishes rather than merely builds. It uses floating Go `stable` and GoReleaser `latest`.
- GoReleaser's `before` hook runs `go mod tidy` but CI does not verify or commit the result, so run `go mod tidy` before tagging. Ensure SQL generation and the focused build are also current before pushing a release tag.
- Releases cross-build Linux, macOS, and Windows on `amd64`/`arm64` with `CGO_ENABLED=0`; archives include `docs/man/doit.1` and `LICENCE`, and the formula is published to `aakku106/homebrew-tap`. The workflow needs `GITHUB_TOKEN` and `HOMEBREW_TAP_TOKEN` with repository/package write access and tap write access.

`LICENCE` is GPLv3. Handwritten Go files carry its header; generated files and `sql/assects.go` do not.
