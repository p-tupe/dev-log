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

## How to

### Read a file line by line

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

> Make invalid states unrepresentable

State-Type pattern dictates that you have a different type for each state of your value. For example, if you have a "Post" type that must be drafted, reviewed, approved and then published (if not rejected), then each state of that type becomes a new type.

To contrast with sometime like Go, where we would have a `status` key inside a mutable struct:

```go
type Post struct {
	status    string
	created_at   time.Time
	inreview_at  time.Time
	approved_at  time.Time
    // ...
}
```

And if we wish to move a post from "draft" to "inreview", we would add a method with a check like so:

```go
func (p *Post) review() (*Post, error) {
	if p.status != "draft" {
		return nil, errors.New("Post not a draft")
	}

	p.status = "inreview"
	p.inreview_at = time.Now()
	return p, nil
}
```

Of course, we can tighten the "state" by having defining a const enum and so on, but the basic flow remains the same.

In Rust, it would idiomatic to do something like so:

```rust

struct Draft {
    content: String,
    created_at: SystemTime,
}

impl Draft {
    fn new(content: String) -> Draft { /*..*/ }
    fn review(self) -> InReview { /*..*/ }
}

struct InReview {
    content: String,
    created_at: SystemTime,
    inreview_at: SystemTime,
}

impl InReview {
    fn approve(self) -> Approved { /*..*/ }
    fn reject(self) -> Rejected { /*..*/ }
}

// ... and so on for other types/impls
```

Why does this matter? For two reasons:

1. You cannot call functions not available on that type
2. Once a valid function is called, the value is "consumed" and no longer available

Rust functions consume their parameters like so:

```rust
fn main() {
    let some_val = vec![1, 2, 3];
    // add_one consumes some_val
    let another_val = add_one(some_val);
    // so some_val is no longer available
    // this code, when uncommented, will not compile
    // println!("{some_val:?}");
    println!("{another_val:?}");
}

fn add_one(val: Vec<i32>) -> Vec<i32> { val.iter().map(|i| i + 1).collect() }
```

And so, for our Post example:

```rust
fn main() {
    let draft = Draft::new(String::from("some content"));

    // Cannot approve before sending for review - no method `approve` on `Draft`
    // draft.approve();
    let inreview = draft.review();

    // Cannot review again by mistake - `draft` has been "consumed"
    // draft.review();
    let approved = inreview.approve();

    // Cannot reject once approve - no method `reject` on approved
    // approved.reject();
    let published = approved.publish();

    println!(
        "{}, {:?}, {:?}, {:?}, {:?}",
        published.content,
        published.created_at,
        published.inreview_at,
        published.approved_at,
        published.published_at
    );

    // Ideal for shadow binding
    let post = Draft::new(String::from("some content"));
    // Or chaining
    let post = post.review().reject();
    println!(
        "{}, {:?}, {:?}, {:?}",
        post.content, post.created_at, post.inreview_at, post.rejected_at
    );
}
```

Try uncommenting the function calls and compiling to see what errors pop up.

And that's it! The State-Type Pattern. It is possible to implement this in other languages as well, but since they don't naturally consume their variables like here, it can get a bit unergonomic and less safe.

Finally, taking this too far in the name of safety means losing a lot on flexibility. In real-world apps, for instance, one might want a Post to be automatically approved if written by an internal team (and skip the "review" phase). It would be trivial to add a conditional in Go's implementation whereas you might have to rethink the whole implementation in our current design.
