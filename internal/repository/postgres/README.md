# PostgreSQL repository

This package will contain sqlc-generated query code and repository adapters.

Rules:

- Pass `context.Context` to database operations.
- Use parameterized SQL; never interpolate user input into query strings.
- Keep transactions explicit for multi-step workflows.
- Translate database errors at the repository boundary.
- Do not expose database driver errors directly to API clients.
- Keep generated files generated; edit SQL sources instead.
