# gostarter

Template for quickly bootstrapping a REST API with Go.

## Prerequisites

- Go v1.22
- Docker

## Get started

1. Create .env file

```
cp .env.example .env
```

2. Start database

```
make db-up
```

3. Run server

```
make run
```

## Features

- MongoDB
- Unit testing
- GitHub Actions
- Linting
- Logging
- Live reloading
- Containerization
- Environment variables
