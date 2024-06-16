# gostarter

## Prerequisites

- Go v1.22.3
- Docker
- goose (`go install github.com/pressly/goose/v3/cmd/goose@latest`)

## Get started

1. Create .env file

```
cp .env.example .env
```

2. Start database

```
make db-start
```

3. Run migrations

```
make migrations-up
```

4. Run server

```
make run
```

## Features

- Postgres database
- Unit testing
- Database migrations
- GitHub Actions
- Linting
- Containerization
- Makefile
- Environment variables

## License
