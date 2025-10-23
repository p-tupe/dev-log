---
modified: Wed Jan  3 21:23:27 EST 2024
---
# Redis

- [Official Site](https://redis.io/)

## Streams

> An append-only log

- [Source](https://redis.io/docs/data-types/streams/)

### Intro

...

### Commands

- XADD: Add a new entry

  - eg `XADD <name> * key1 value1 key2 value2 ...`

- XREAD: Read entries from a position moving forward

  - eg `XREAD `

- XRANGE: Read entries from x to y

  - eg `XRANGE <name> <from>`

- XLEN: Length of a stream

  - eg `XLEN <name>`
