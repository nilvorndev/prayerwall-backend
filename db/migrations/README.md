# Database migrations

Use versioned SQL migrations with Goose after the existing Django schema has been inspected.

**Do not run migrations against production during foundation work.** First restore a backup into a separate development database and document current tables, constraints, indexes, and row counts. The initial migration for this rewrite should be additive and reviewed against that restored copy.

Migration naming convention: `YYYYMMDDHHMMSS_description.sql`.
