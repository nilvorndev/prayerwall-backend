# PrayerWall Backend

Go backend rewrite for PrayerWall.

## Foundation stack

- Go
- Gin
- PostgreSQL + pgx
- sqlc for generated query access
- Goose for versioned SQL migrations
- Docker Compose for local development

## Prerequisites

- Go toolchain matching `go.mod`
- Docker Engine and Docker Compose (for container-based development)

## Local development

1. Copy `.env.example` to `.env` if you want to run the API directly on the host.
2. Start the API and local PostgreSQL using Docker Compose:

   ```sh
   docker compose -f deploy/compose.yaml up --build
   ```

3. Verify health:

   ```sh
   curl -i http://localhost:8080/healthz
   curl -i http://localhost:8080/readyz
   ```

4. Run checks:

   ```sh
   make ci
   ```

## Endpoints

- `GET /healthz` — process health; does not depend on PostgreSQL.
- `GET /readyz` — readiness; checks PostgreSQL connectivity.

## Safety rules

- Never commit `.env` or production credentials.
- Never point this starter at the production database.
- Do not create or apply schema-changing migrations until the existing Django schema has been inspected and a backup restore has been verified.
- Keep the existing Django service authoritative until a planned cutover.
- Use a single authoritative writer during production migration.

## Next steps

1. Inventory the existing Django API and PostgreSQL schema.
2. Add compatibility tests and the OpenAPI contract.
3. Configure sqlc and Goose against a restored development database.
4. Implement identity workflows, then prayer workflows.
