---
modified: Thu Oct 20 15:32:23 EDT 2022
---
# Postgresql

- [Official site](https://www.postgresql.org)

- [Official docs](https://www.postgresql.org/docs)

- [Tutorial](https://www.postgresqltutorial.com/)

- [Quick Reference](https://zaiste.net/posts/postgresql-primer-for-busy-people/)

## Notes

- Enter psql shell: `[<docker>] psql -U <user> -W -d <database>`

- Connection info: `# \c [<db>] [<user>]`

- List all databases: `# \l`

- List all tables: `# \d`

- List all users: `# \du`

- Backing up postgres (use `pg_dumpall` for... all)

  ```sh
  pg_dump -U <user> <database> | xz -4 > /tmp/postgres_dump_$(date +"%Y-%m-%d_%H_%M").sql.xz
  ```

- Restoring from dump

  ```sh
  xz -d < /tmp/postgres_dump_<date>.sql.xz | <docker exec> psql -U  <user> -d <database>
  ```

  - If restoring a particularly big dump or throwing errors, `docker cp` it inside container and restore from within.

- To run above commands inside docker container (denoted by `<docker exec>`)

  ```sh
  docker exec -it [-u <user>] <postgres-container-name> <command>
  ```

### Setup for new Apps

> Assuming a postgres instance is already running and a new app is being setup that requires a database in it.

1. Create a new role (<user>/password)
1. Create a new database for app
1. Restrict acess for the role to that database

```sql
# CREATE DATABASE <database>;
# CREATE USER <user> WITH ENCRYPTED PASSWORD '<pass>';
# ALTER DATABASE <database> OWNER TO <user>;
# GRANT ALL PRIVILEGES ON DATABASE <database> TO <user>;
```

> Use this connection string for access: `postgres://<user>:<pass>@<host>:<port>/<database>`
