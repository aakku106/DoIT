<div align="center">

# DoIT

**A sleek CLI todo/task manager built with Go & WASM SQLite3**

[![License: GPLv3](https://img.shields.io/badge/License-GPLv3-blue.svg)](LICENCE)
[![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8.svg)](go.mod)

</div>

> **Created by** Adarasha Gaihre (aakku106) on 2026-09-02 00:15 NPT (+0545)
> **Last updated by** Adarasha Gaihre (aakku106) on 2026-09-25 23:45 NPT (+0545)

DoIT is a fast, local-first command-line task manager. Tasks live in four
lists — **todos**, **completed**, **trash**, and **ignored** — and everything
is stored in a plain SQLite database, so you can stay keyboard-only and your
data stays on your machine.

## Features

- **Local storage** — data lives in a `.doit/doit.db` SQLite database (WAL journal mode).
- **Four task lists** — draft, complete, trash, and ignore your tasks.
- **Multi-task input** — add several tasks in one `doit add` invocation, comma-separated.
- **Move tasks between lists** — recover or reclassify with a single command.
- **Interactive confirmation** — destructive remove/clear actions ask before running; bulk `nuke` commands require a second typed phrase.
- **Cross-platform** — prebuilt binaries for Linux, macOS, and Windows (amd64 + arm64).

## Installation

### Homebrew (macOS / Linux)

```sh
brew tap aakku106/homebrew-tap
brew install doit
```

This uses the third-party `aakku106/homebrew-tap` repository. Review the
formula before installing if you prefer to inspect third-party taps — you can
see it at [`aakku106/homebrew-tap`](https://github.com/aakku106/homebrew-tap).

### From source

Requires [Go 1.25+](https://go.dev/dl/). The module root holds no Go package, so
clone and build rather than using `go install`:

```sh
git clone https://github.com/aakku106/DoIT.git
cd DoIT
go build -o doit ./cmd/todo
```

### From release archives

Download the archive matching your OS/arch from the
[releases](https://github.com/aakku106/DoIT/releases) page, extract it, and put
the `doit` binary on your `PATH`.

## Getting Started

DoIT stores its database in a `.doit/` directory. Initialize one in (or above)
your current directory, then start adding tasks:

```sh
# Start a new project
doit init

# Add and list tasks
doit add "buy milk"
doit add "call mom"
doit list
```

> DoIT finds the nearest `.doit/` by walking **upward** from your current
> directory, so you can use it from any subdirectory of an initialized project.

## Usage

### Core

| Command | Description |
| --- | --- |
| `doit init` | Create a `.doit/` repository + database |
| `doit add <task>` | Add a new todo (one task, or several — see below) |
| `doit list` (`ls`) | List active todos |
| `doit done <id>` (`d`) | Mark a todo done (→ completed) |
| `doit remove <id>` (`rm`) | Move a todo to trash (confirmation) |
| `doit nuke` (`n`) | Permanently clear the **todo** list (two-step confirmation) |

> **IDs are display indices, not database row IDs.** Use the number shown next
> to a task in `list` output. (`doit add` does print a database row ID — that
> one is *not* usable with the commands above.)

### Adding several tasks at once

Passing more than one argument switches `add` to multi-task mode. Separate the
tasks with a comma that is its own argument (surrounded by spaces, or quoted on
its own), and optionally give each one a deadline with `-t=`:

```sh
doit add taskA -t=2h , "task B" -t=1mo
```

Tasks are added in the order given. A comma attached to a task
(`"taskA,"`) is treated as part of that task's text and will be rejected.

### List management

Each of `completed` (`c`), `trash` (`t`), and `ignored` (`i`) accepts a subcommand:

```sh
doit <completed|trash|ignored> <list|remove|nuke>
doit c list        # show completed tasks
doit t list        # show trashed tasks
doit t rm 0        # permanently delete trashed task 0
doit c nuke        # permanently clear all completed (confirmation)
```

### Clearing a list

Every `nuke` — top-level or per-table — uses a two-step confirmation. Answer the
`(Y/N)?` prompt with a strict uppercase `Y`, then type the exact phrase:

| Command | Confirmation phrase |
| --- | --- |
| `doit nuke` | `YeS NuKe ToDos` |
| `doit c nuke` | `YeS NuKe CoMpleteD` |
| `doit t nuke` | `YeS NuKe TrAsH` |
| `doit i nuke` | `YeS NuKe IgNoreD` |

### Moving tasks between lists

```sh
doit mv <t|c|i> <id> [<t|c|i>]
```

The source is chosen by the first `<t|c|i>` argument and `id` is the 0-based
index within that list. With no target, the task moves back to **todos**:

```sh
doit mv t 0        # move trashed task 0 back to todos
doit mv t 0 c      # move trashed task 0 to completed
doit mv c 2 i      # move completed task 2 to ignored
```

## Known limitations

- `doit t rm <id>` currently deletes from the **completed** table instead of the
  trash table. Use `doit t nuke` to empty the trash until this is fixed.
- An id equal to the length of a list is not rejected, and can panic or affect
  the wrong task.
- `doit d <non-numeric-id>` warns but still acts on the first task.
- The `-t=` deadline is validated and then discarded; nothing is stored and no
  task ever expires.
- Malformed multi-task input (leading/trailing comma, a second task argument
  without `-t=`) aborts with a panic rather than a clean error message.
- Sessions are not implemented; the session name is hardcoded to `todo`.

`man doit` documents all of these in its `BUGS` section.

## Building from Source

```sh
git clone https://github.com/aakku106/DoIT.git
cd DoIT

go build -o doit ./cmd/todo
```

Run without building first:

```sh
go run ./cmd/todo add "wash car"
go run ./cmd/todo list
```

### Dev build (`xdoit`)

The repo includes a `dev-build.sh` helper that builds the binary and installs
it to `~/go/bin/xdoit` (note the `x` prefix, so it won't shadow a real
`doit` install):

```sh
./dev-build.sh
xdoit add "test task"
```

> `xdoit` is a **development/testing build**. It is not versioned or
> officially distributed, may be unstable, and is not recommended for
> production use — use it after cloning the repo to verify the app works, or
> simply use `go build` / `doit` from a release instead.

### Manpage

A man page is included at `docs/man/doit.1` and shipped in release archives and
via Homebrew. View it with `man doit`.

## Documentation

- [Project overview](docs/project-overview.md) — architecture and file-by-file details
- [`docs/man/doit.1`](docs/man/doit.1) — manpage

## Contributing

Contributions are welcome! This is a small project, so please open an issue or
PR for any improvement. If you change the SQL schema or queries, regenerate the
store layer:

```sh
# after editing sql/schema.sql or sql/queries.sql
sqlc generate
```

> Never hand-edit files under `internal/store/` — they are generated by `sqlc`.

## License

GPLv3 — see [LICENCE](LICENCE).

Copyright (C) 2026 Adarasha Gaihre (aakku106).
