# go-to-do

[![Go](https://github.com/denzasikora-lab/go-to-do/actions/workflows/go.yml/badge.svg)](https://github.com/denzasikora-lab/go-to-do/actions/workflows/go.yml)

Telegram **to-do bot** written in Go: Postgres-backed tasks, **FSM** conversational flows, and **inline keyboards** with semantic button colors ([Telegram Bot API `style` field](https://core.telegram.org/bots/api#inlinekeyboardbutton)—requires recent clients).

## Features

- **Task lifecycle**: create (title → description → priority), list with filters, inspect, mark done / reopen, edit title & description, delete with confirmation.
- **FSM + persistence**: `bot_sessions` stores state and JSON payload; tasks live in `todos`.
- **Docker Compose**: Postgres + bot container (see `docker-compose.yml`).
- **Layered layout**: `domain` → `repository` → `bot`, small focused files.

## Requirements

- Go **1.22+**
- PostgreSQL **16+** (or use Compose)
- Telegram bot token from [@BotFather](https://t.me/BotFather)

## Quick start (local)

```bash
cp .env.example .env
# Set TELEGRAM_BOT_TOKEN and POSTGRES_DSN (e.g. postgres://user:pass@127.0.0.1:5432/dbname?sslmode=disable)

go run ./cmd/bot
```

## Docker Compose

```bash
cp .env.example .env
# Set TELEGRAM_BOT_TOKEN. Default POSTGRES_DSN targets the `postgres` service hostname.

docker compose up --build
```

Ensure a Docker engine is running (Docker Desktop, Colima, etc.).

## Environment

| Variable | Description |
|----------|-------------|
| `TELEGRAM_BOT_TOKEN` | **Required**. Bot API token. |
| `POSTGRES_DSN` | **Required**. `postgres://user:password@host:port/database?sslmode=disable` |

## Repository

- **GitHub**: [github.com/denzasikora-lab/go-to-do](https://github.com/denzasikora-lab/go-to-do)

### Suggested repository metadata (GitHub → Settings → General)

- **Description**: `Telegram to-do bot in Go — Postgres, FSM, colored inline keyboards, Docker.`
- **Topics**: `go`, `telegram-bot`, `postgresql`, `docker`, `fsm`, `todo`

## License

MIT — see [LICENSE](LICENSE).
