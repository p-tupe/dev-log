---
modified: Fri Oct 9 16:05:16 2026 -0400
---

# Rust

## Resources

- https://doc.rust-lang.org <- Has many things
- https://rust-unofficial.github.io/patterns/intro.html <- idiomatic rust
- https://blessed.rs/crates | https://lib.rs/ <- Popular crates
- https://rust-lang-nursery.github.io/rust-cookbook/ <- How to use 'em
- https://github.com/rust-unofficial/awesome-rust <- Stuff made in rust
- https://github.com/pretzelhammer/rust-blog <- Good stuff

- https://benjamingeer.eu/en/post/hypermedia-rust/

## Must Know Crates

| name                   | function               |
| ---------------------- | ---------------------- |
| tokio +full            | async runtime          |
| serde +derive          | (de)serialization glue |
| serde_json             | json parser            |
| quick_xml +serialize   | xml parser             |
| tracing (& subscriber) | logging in axum        |
| log (& env_logger)     | logging in general     |
| jiff                   | datetime stuff         |
| axum                   | web server             |
| regex                  | regular expressions    |
| rand                   | random numbers         |
| anyhow                 | error handling         |
| confy                  | app configuration      |
| dirs                   | more user directories  |
| walk_dir               | to walk directories    |
| sqlx                   | sqlite, postgres       |
| reqwest/ureq +json     | http libs              |
| clap                   | cli arg parsing        |
| indicatif              | cli progress bars      |

## How to

### Conquer Command

```rust
Command::new(x).spawn()?; //  fire 'n forget
Command::new(x).output()?; // fire 'n ... wait (includes status along with rest)
Command::new(x).status()?; // wait until complete to see
```

### Get `Date.now()` JS equivalent

```rust
let now = std::time::SystemTime::now().duration_since(SystemTime::UNIX_EPOCH)?.as_millis();
```

### Optimize release binary

In the project's `cargo.toml`:

```toml
[profile.release]
# remove symbol names (and DWARF) from the binary — smaller, but backtraces lose function names [default: "none"]
strip = "symbols"
# optimize across crates, not just within one — "thin" gets most of "fat"'s gain for far less link time [default: false]
lto = "thin"
# compile the crate as a single unit so LLVM sees the whole thing — better code, slower build [default: 16]
codegen-units = 1
# panic terminates instead of unwinding the stack — smaller/faster, but no catch_unwind and no unwind cleanup [default: "unwind"]
panic = "abort"
```

### Embed a file into the binary

```rust
let embedded_file = include_str!("./path/to/file");
```

Can use this to add README as doc string

```rust
#![doc = include_str!("../README.md")]
```

### Build a cross-platform executable

> As seen on [official docker rust image](https://hub.docker.com/_/rust/#compile-your-app-inside-the-docker-container)

```bash
# change os and arch as required
docker run --rm --platform os/arch --user "$(id -u)":"$(id -g)"  \
    -v "$PWD":/usr/src/app -w /usr/src/app \
    rust:slim cargo build --release
# use rust:alpine for static musl'd bin
```

> See [cross](https://github.com/cross-rs/cross) for complex usecases (does not work on MacOS though :/)

### Run rust from a (minimal) docker file

- [github.com code-example](https://github.com/p-tupe/code-examples/tree/main/rust/with-docker)

### Run rust as standalone script

```bash
#!/usr/bin/env -S cargo +nightly -Zscript
---cargo
[dependencies]
serde_json = "*"
---
fn main() {}
```

### Quickly convert a digit (0-9) into char

```rust
(digit + b'0') as char
```

### Buffer a file by line

> https://doc.rust-lang.org/stable/rust-by-example/std_misc/file/read_lines.html

```rust
let file = File::open("./cargo.toml")?;
for line in BufReader::new(file).lines() {
    println!("{}", line?);
}

BufReader::new(File::open("./cargo.toml")?)
    .lines()
    .map_while(Result::ok)
    .for_each(|l| println!("{}", l));
```

## State-Type Pattern

- [priteshtupe.com/posts/rust-actors/](https://priteshtupe.com/posts/state-type-rust/)

### Actor Pattern

- [priteshtupe.com/posts/rust-actors/](https://priteshtupe.com/posts/rust-actors/)
