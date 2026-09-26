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
| tokio                  | async runtime           |
| clap                   | cli arg parsing         |
| serde (+ others)       | struct to json/toml/etc |
| tracing (+ subscriber) | logging                 |
| time                   | datetime stuff          |
| axum                   | web server              |
| regex                  | regular expressions     |
| rand                   | random numbers          |
| anyhow                 | error handling          |
| config                 | configuration file      |
| dirs                   | user directories        |
| walk_dir               | to walk directories     |
| rusqlite               | sqlite interface        |

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

### Embed a file into a binary

```rust
let embedded_file = include_str!("./path/to/file");
```

### Run rust as standalone script

```bash
#!/usr/bin/env -S cargo +nightly -Zscript
---cargo
[dependencies]
serde_json = "*"
---
fn main() { println("this is a script!"); }
```

### Quickly convert a digit (0-9) into char

```rust
(digit + b'0') as char
```

### Buffer a file by line

> https://doc.rust-lang.org/stable/rust-by-example/std_misc/file/read_lines.html

```rust
use std::{
    fs::File,
    io::{BufRead, BufReader, Result},
};

fn main() -> Result<()> {
    let file = File::open("./cargo.toml")?;
    let line = String::new();
    for line in BufReader::new(file).lines() {
        println!("{}", line?);
    }
    Ok(())
}

```

## State-Type Pattern

- [priteshtupe.com/posts/rust-actors/](https://priteshtupe.com/posts/state-type-rust/)

### Actor Pattern

- [priteshtupe.com/posts/rust-actors/](https://priteshtupe.com/posts/rust-actors/)
