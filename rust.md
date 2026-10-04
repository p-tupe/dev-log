---
modified: "Sat May  9 10:26:00 EDT 2026"
---

# Rust

## Resources

- https://doc.rust-lang.org <- Has many things
- https://rust-unofficial.github.io/patterns/intro.html <- idiomatic rust
- https://blessed.rs/crates | https://lib.rs/ <- Popular crates
- https://rust-lang-nursery.github.io/rust-cookbook/ <- How to use 'em
- https://github.com/rust-unofficial/awesome-rust <- Stuff made in rust
- https://github.com/pretzelhammer/rust-blog <- Good stuff

## Must Know Crates

| name                   | function                |
| ---------------------- | ----------------------- |
| tokio +full            | async runtime           |
| clap                   | cli arg parsing         |
| serde +derive          | struct to json/toml/etc |
| tracing (& subscriber) | logging in axum         |
| log (& env_logger)     | logging in general      |
| time                   | datetime stuff          |
| axum                   | web server              |
| regex                  | regular expressions     |
| rand                   | random numbers          |
| anyhow                 | error handling          |
| config                 | configuration file      |
| dirs                   | user directories        |
| walk_dir               | to walk directories     |
| rusqlite               | basic sqlite            |
| sqlx                   | async sqlite            |

## Crash course on Result/Option handling nomenclature

```rust
// Assume op on Result<Ok(Value), Err(Error)> or Option<Some(Value) | None>

// ok*

// unwrap*

// *and*
// *or*

// *_else
// *_then

// map*
// filter*
// reduce*
```

### Discard None (Optional) values in a loop

```rust
let x = [Some(1), None, Some(2), None, Some(3)];

// Using let Some(x) = y else { continue }
let mut sum = 0;
for v in x {
    let Some(v) = v else { continue };
    sum += v;
}

// Using filter_map
let fsum = x.iter().filter_map(|v| v.map(|v| v)).fold(0, |s, e| s + e);
```

### Chose this or that (Optional) if they exists, else do something else

```rust
// .or_else is the main part
let Some(editor) = env::var_os("VISUAL").or_else(|| env::var_os("EDITOR")) else {
    println!("no editor found");
};

let _ = Command::new(editor).arg(path).status()?;
```

## How to

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

> See [cross](https://github.com/cross-rs/cross) for complex usecases

```bash
# change os and arch as required
docker run --platform os/arch --rm \
    -v "$PWD":/usr/src/app -w /usr/src/app \
    rust:latest cargo build --release
```

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
