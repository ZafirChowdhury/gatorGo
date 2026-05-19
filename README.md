# gatorGo

A command-line RSS feed aggregator built in Go.

---

## What it does

- Register and switch between users
- Add RSS feed URLs and follow/unfollow them
- Run a background aggregator that periodically fetches new posts
- Browse your followed feeds' latest posts directly in the terminal

---

## Getting started

### Prerequisites

- [Go](https://go.dev/dl/) (1.21+)
- [PostgreSQL](https://www.postgresql.org/download/) running locally

### Install

```bash
go install github.com/ZafirChowdhury/gatorGo@latest
```

### Config

Create a config file at `~/.gatorconfig.json`:

```json
{
  "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable",
  "current_user_name": ""
}
```

Replace `username`, `password`, and the database name as needed. Create the database in psql first:

```sql
CREATE DATABASE gator;
```

### Run migrations

The project uses `goose` for migrations. From the repo root:

```bash
goose -dir sql/schema postgres "your-connection-string" up
```

---

## Commands

```bash
# Create a user and set them as current
gator register alice

# Switch users
gator login alice

# Add an RSS feed (auto-follows it too)
gator addfeed "Hacker News" https://hnrss.org/frontpage

# Follow an existing feed by URL
gator follow https://hnrss.org/frontpage

# Start the aggregator (fetches every 30s, 1m, etc.)
gator agg 30s

# Browse your latest posts (default 2, or pass a limit)
gator browse 5

# See all feeds in the system
gator feeds

# See what you're following
gator following

# List all users
gator users
```

---

## What I learned

* Structuring a CLI in Go — registering named commands in a map and dispatching to handlers at runtime
* Middleware pattern in Go — wrapping handlers with `middlewareLoggedIn` to inject auth without touching every handler
* Using `sqlc` to generate type-safe Go from raw SQL, keeping the database layer explicit with no ORM
* Writing and running PostgreSQL migrations with `goose`
* Periodic background work with `time.NewTicker` and a blocking `for` loop
* Handling nullable SQL columns in Go with `sql.NullString` and `sql.NullTime`
* Fetching and parsing RSS/XML feeds over HTTP using `encoding/xml`
* Deduplicating database writes by catching unique-constraint errors instead of pre-checking
