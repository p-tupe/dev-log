---
modified: "Sat Apr 25 13:32:07 EDT 2026"
---

# Parallel Log Cruncher — Language-Agnostic Learning Spec

> Build this in any language you're learning.

---

## Mandatory

### 1. Parse a log file

```
timestamp level ip path method status duration
2026-04-23T14:22:01Z INFO 192.168.1.10 /api/users GET 200 15ms
```

**What to do:** Read the file one line at a time. Split each line by whitespace. Validate each field:

- `timestamp` — ISO 8601 format (parseable datetime)
- `level` — one of INFO, WARN, ERROR, DEBUG
- `ip` — dotted decimal IPv4 (4 octets, 0-255 each)
- `path` — starts with `/`
- `method` — one of GET, POST, PUT, DELETE, PATCH
- `status` — integer 100-599
- `duration` — ends with `ms`, parse to integer

Skip any line that doesn't match. Print a warning to stderr for each skipped line. Keep count of valid vs skipped.

**Line-by-line matters** — don't read the whole file into memory at once. Real log files can be gigabytes. Use the language's buffered reader.

### 2. Compute analytics

From the parsed lines, compute:

- **Errors per endpoint** — group by `path`, count how many have `status >= 400`
- **P95 latency per endpoint** — for each `path`, collect all `duration` values, sort them, pick the value at the 95th percentile index (index = ceil(0.95 × count) - 1)
- **Top 10 IPs by request count** — count requests per IP, sort descending, take first 10

### 3. Output

Write `report.json` containing:

```json
{
  "total_requests": 15000,
  "skipped_lines": 3,
  "errors_per_endpoint": { "/api/users": 12, "/api/checkout": 45 },
  "p95_latency_ms": { "/api/users": 120, "/api/checkout": 890 },
  "top_ips": [
    { "ip": "192.168.1.10", "count": 2300 },
    { "ip": "10.0.0.5", "count": 1800 }
  ]
}
```

Use the language's standard JSON serializer. Don't build strings manually.

### 4. CLI

```
logcrunch <file>
```

- `<file>` — path to log file (required positional argument)

The tool runs the analytics and writes report.json.

---

## Optional

| Option                          | What you learn                                                                              |
| ------------------------------- | ------------------------------------------------------------------------------------------- |
| **Geolocation + rate limiting** | Async HTTP client, bounded concurrency, rate limiter, shared state, graceful error handling |
| **HTTP server**                 | Web framework, routing, JSON responses, graceful shutdown                                   |
| **Benchmark mode**              | Timing instrumentation, reporting, separating measurement from logic                        |
| **Unit tests**                  | Test structure, fixtures, assertion patterns, code coverage                                 |
| **Streaming parse (stdin)**     | Reading from pipes, handling unbounded input, testable with `echo ... \| logcrunch`         |

### Geolocation + rate limiting

For each **unique IP**, fetch its country code:

```
GET http://ip-api.com/json/{ip}
{"countryCode": "US", ...}
```

**Constraints:**

- Max 45 requests per minute (API rate limit)
- Handle timeouts, connection failures, and non-200 responses
- Mark failed lookups as `"XX"` (unknown)
- Cache results — don't re-request the same IP

**How:** Collect unique IPs from parsed lines. Spawn `--workers` concurrent tasks. Each worker acquires a permit from a rate limiter before making the request. Write results to a shared map protected by the language's concurrency primitive (mutex, channel, actor).

**If the API is down:** The tool should still complete — mark all unknown, print a warning, continue.

### HTTP server

After processing the file, start a web server on the port specified by `--port` (default 8080).

Routes:

- `GET /stats` — full JSON analytics (same as report.json)
- `GET /healthz` — `{"status": "ok"}` + 200

**How:** Start the server only after all processing is done. Bind to `0.0.0.0`. Handle SIGINT/SIGTERM for graceful shutdown.

### Benchmark mode

`--bench` flag: process the file, print timing breakdown to stdout, then exit. Do NOT start the HTTP server. Do NOT write report.json.

Output:

```
parse: 1.24s
analytics: 0.08s
total: 1.32s
lines: 15000
parse_rate: 12096 lines/s
memory_peak: 24.2 MB
```

**How:** Wrap each phase with a high-resolution timer. Report memory if the language provides it (otherwise skip).

### Unit tests

Test the parser in isolation — feed it known-good and known-bad lines, verify it returns the right struct or error.

Test analytics — feed it a parsed dataset with known values, verify P95, error counts, and top IPs match.

Test output — parse a small file, run analytics, check that report.json has the expected schema and values.

**How:** Use the language's standard test framework. Put test data inline (small strings) or in a `testdata/` directory. Test edge cases: empty file, all errors, single line, malformed fields.

### Streaming parse (stdin)

Accept input from stdin when no file argument is given:

```
cat access.log | logcrunch
```

**How:** If `<file>` is missing or `-`, read from stdin. Same line-by-line loop, different source.

---

## Concepts

| Concept                                 | Why it matters                                                                                                                   |
| --------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| **Buffered I/O**                        | Reading line by line without loading the whole file into memory. Every language has a `BufReader` or equivalent.                 |
| **String splitting / pattern matching** | Parsing structured text. Some languages use regex, some use destructuring, some use split-and-index. The trade-offs matter.      |
| **Structs / records / data classes**    | Defining typed data. How the language models structured data affects everything downstream.                                      |
| **Sum types / enums**                   | Level (INFO, WARN, ERROR, DEBUG) is a natural enum. Does the language have real sum types or do you fake it with strings?        |
| **Serialization (JSON)**                | Converting memory to wire format. Derive macros vs reflection vs manual.                                                         |
| **Error handling**                      | What happens when a line is malformed, a file doesn't exist, or a network request fails. Idiomatic vs try-catch vs Result types. |
| **Concurrency primitives**              | Threads, async/await, goroutines, fibers — how the language models parallel work.                                                |
| **Rate limiting**                       | A real-world constraint that forces you to think about time and resource fairness.                                               |
| **Shared mutable state**                | Multiple tasks writing to the same map. Mutex? Channels? Actors? STM?                                                            |
| **HTTP client / server**                | Making and serving requests. Timeouts, connection pooling, middleware.                                                           |
| **CLI argument parsing**                | Libraries that generate help text, validate types, and handle missing arguments.                                                 |
| **Percentile calculation**              | Sorting + indexing. Straightforward but easy to get wrong (off-by-one on the index).                                             |
| **Graceful shutdown**                   | Catching SIGINT, finishing in-flight work, flushing output, then exiting cleanly.                                                |

---

## Resources

### Buffered I/O

- https://en.wikipedia.org/wiki/Buffered_reader
- https://stackoverflow.com/questions/7867537/what-is-buffered-i-o-and-how-does-it-improve-performance

### JSON serialization

- https://www.json.org/json-en.html
- https://en.wikipedia.org/wiki/Serialization

### Async I/O

- https://en.wikipedia.org/wiki/Asynchronous_I/O
- https://journal.stuffwithstuff.com/2015/02/01/what-color-is-your-function/

### Rate limiting patterns

- https://en.wikipedia.org/wiki/Rate_limiting
- https://www.cloudflare.com/learning/bots/what-is-rate-limiting/

### Worker pool pattern

- https://en.wikipedia.org/wiki/Thread_pool
- https://go.dev/doc/faq#goroutines (approachable explanation of the pattern even if you don't use Go)

### Concurrency vs parallelism

- https://go.dev/blog/waza-talk (Rob Pike's talk — language agnostic, diagrams)
- https://en.wikipedia.org/wiki/Concurrency_(computer_science)

### Shared mutable state

- https://en.wikipedia.org/wiki/Shared_memory_(interprocess_communication)
- https://en.wikipedia.org/wiki/Mutual_exclusion

### Graceful shutdown

- https://en.wikipedia.org/wiki/Graceful_degradation
- https://learn.microsoft.com/en-us/dotnet/standard/garbage-collection/background-gc (GC-specific but the signal-handling pattern is universal)

### Percentile calculation

- https://en.wikipedia.org/wiki/Percentile
- https://www.statisticshowto.com/probability-and-statistics/percentiles-rank-range/

### Signal handling

- https://en.wikipedia.org/wiki/Unix_signal
- https://www.gnu.org/software/libc/manual/html_node/Basic-Signals.html

### Testing philosophy

- https://en.wikipedia.org/wiki/Unit_testing
- https://en.wikipedia.org/wiki/Test-driven_development
- https://www.youtube.com/watch?v=EZ05e7akOLU (Ian Cooper — TDD, where it fits)

### CLI argument parsing patterns (general)

- https://en.wikipedia.org/wiki/Command-line_interface
- https://pragprog.com/titles/dccar2/build-awesome-command-line-applications-in-ruby/ (practical patterns, applies to any language)

---

All resources are about concepts, not language APIs. If you need the API docs for a specific feature, search: `"<language>" "<concept>"`.

---

## Quick log generator (bash)

```bash
i=0; s=42; l=(INFO WARN ERROR DEBUG); m=(GET POST PUT DELETE PATCH); p=(/api/users /api/checkout /api/login /api/products /api/search /api/orders /api/reviews)
RANDOM=$s; while ((i++ < ${1:-1000})); do
  t=$(date -u -d "2026-04-23T00:00:00+00:00 + $i seconds" +%FT%TZ 2>/dev/null || date -u -j -v+${i}S -f "%Y-%m-%dT%H:%M:%S%z" "2026-04-23T00:00:00+0000" +%FT%TZ 2>/dev/null || printf "2026-04-23T%02d:%02d:%02dZ" $((i/3600)) $(((i%3600)/60)) $((i%60)))
  st=200; ((RANDOM > 28000)) && st=$((400+RANDOM%100))
  echo "$t ${l[$RANDOM%4]} $((RANDOM%256)).$((RANDOM%256)).$((RANDOM%256)).$((RANDOM%256)) ${p[$RANDOM%7]} ${m[$RANDOM%5]} $st $((RANDOM%500+1))ms"
done > test.log
```
