# How to use `pgxgen` tool

This repository explain how to use [pgxgen tool](https://github.com/tkcrm/pgxgen)

## Local development

### Install tools

```bash
make install-tools
```

### Enviroments

Create `.env` file with your enviroments in root repository based on `.env.example` file

### Create database

```bash
createdb pgxgenexample
```

### Make migrations

```bash
make migrate up
```

### Start microservice

```bash
make start
```

### Create new migration

```bash
// replace MIGRATION_NAME.
// Example: make db-craete-miration book_rating
make db-craete-miration MIGRATION_NAME
```

### Generate crud sql and store

- Apply new migrations
- Edit `pgxgen.yaml` in root repository. Documentation available [here](https://github.com/tkcrm/pgxgen#configure-pgxgen)
- Execute command `make gensql`

> To override types use [this documentation](https://docs.sqlc.dev/en/stable/reference/config.html?highlight=override#per-column-type-overrides)
