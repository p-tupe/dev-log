---
modified: Mon Nov 14 23:06:32 EST 2022
---
# Sqlite 3

## Cli shell commands

- Help

  ```sql
  .help
  ```

- Some useful commands:

  - Open a new connection `.open`
  - List of databases attached `.databases`
  - List of tables present `.tables`
  - Read an sql file `.read <file>`
  - Describe a table `.schema <table>` `pragma table_info("<table>")`
  - Display options `.header on` `.mode column`

- Create table (don't forget semicolon at end)

  ```sql
  create table [if not exists] <name> ( ... );
  ```

- Storage Classes (NOT Data Types!)

  - Null
  - Integer
  - Real
  - Text
  - Blob

- Foreign key (inside create table)

  ```sql
  foreign key ("column_id") references "foreign_table"("foreign_id")
  ```

- Insert Query

  ```sql
  insert into <table> (<col1>, <col2>, ...) values (<v1>, <v2>, ...), (<vk>, <vl>, ...);
  ```

- Update Query

  ```sql
  update <table> set <key> = <val>, ... [from select ...] [where <key> = <val>];
  ```

- Delete Query

  ```sql
  delete from <table> [where <col> = <val>];
  ```

- Select Query

  ```sql
  select [<col1>, <col2> | *]
    from <table1>, [<table2> ...]
    [where <cond1> [<cond2>]]
    [order by <colx> [desc | asc] ...]
    [group by <coly>]
  ```

  - Distinct Query

    ```sql
    select distinct <col1>, <col2> from <table>;
    ```

  - Limit Query

    ```sql
    select ... LIMIT <number>;
    ```

  - Offset Query

    ```sql
    select ... OFFSET <number>;
    ```

    - Offset and Limit may be used together (with order by) for paginated results

  - Like Query

    ```sql
    select * from <table> where LIKE <pattern>;
    ```

    - `<pattern>` is _like_ regex but not really. Here's a comparison:
      - `x%` = `^x+` = starts with "x"
      - `%x` = `.*x$` = ends with "x"
      - `%xyz%` = `.*xyz.*` = contains "xyz"
      - `_x%` = `^.x.*` = has x in second position, `_` is like `.` - means any
      - `x%y` = `^x.*y$` = starts with "x" and ends with "y"

  - Sum | Min | Max | Average | Count - Aggregate functions

  ```sql
  select [SUM | MIN | MAX | AVG | COUNT](<col>) from <table> [group by <table>]
  ```

  - HAVING: Like "WHERE" filter for aggregate functions

  ```sql
  select <col> from <table> [group by <col>] HAVING <[sum | min | max | avg | count ] bool expr>;
  ```

- Backup Sqlite3 (non-interactively)

  ```shell
  $ sqlite3 <name>.db ".backup '<name>.backup.db'"
  ```

  - Alternatively, compress it with `xz`

    ```shell
    $ xz --compress --threads=0 --quiet <name>.backup.db
    ```
