---
modified: "Sat Dec 13 11:19:29 EST 2025"
---

# Go

## Resources

- [Official Site](https://golang.org)

- [Official Learning Resource](https://go.dev/learn/)

- [Go101: Beginners' Book](https://go101.org/)

- [Golang Tutorial: YouTube video tutorial](https://www.youtube.com/watch?v=YzLrWHZa-Kc)

- [How to Write Go Code](https://go.dev/doc/code)

- [Effective Go](https://go.dev/doc/effective_go)

- [Go By Example](https://gobyexample.com/)

- [Awesome Go](https://github.com/avelino/awesome-go)

- [learning-go-by-examples-introduction-448n](https://dev.to/aurelievache/learning-go-by-examples-introduction-448n)

- [learn-golang-series](https://golangbot.com/learn-golang-series)

- [GoBooks](https://github.com/dariubs/GoBooks)

- [goperf.dev](https://goperf.dev)

- [go-concurrency](https://antonz.org/go-concurrency)

- [GolangTraining](https://github.com/GoesToEleven/GolangTraining)

## Notes

- workspace (mono-repos) > module (dir with `go.mod` file) > package (related `.go` files)

- Builtin functions: [golang.org/pkg/builtin](https://golang.org/pkg/builtin)

## How to

### Read from console/terminal

```go
ip := bufio.NewReader(os.Stdin)
fmt.Print("Enter something: ")
name, _ := ip.ReadString('\n')
fmt.Println("You entered: ", name)
```

### Use the regexp package

> [pkg.go.dev/regexp](https://pkg.go.dev/regexp)

```go
// Quick match
matched, err := regexp.MatchString(`foo.*`, "seafood")
fmt.Println(matched, err)

// Find a string
re := regexp.MustCompile(`foo.?`)
fmt.Printf("%q\n", re.FindString("seafood fool"))

// Find capture groups
re := regexp.MustCompile(`a(x*)b`)
fmt.Printf("%q\n", re.FindAllStringSubmatch("-ab-", -1))
fmt.Printf("%q\n", re.FindAllStringSubmatch("-axxb-", -1))

// Replace A with X, B with Y, C with Z
replaceMap := map[string]string{"A": "X", "B": "Y", "C": "Z"}
re := regexp.MustCompile(`(A|C|G|T)`)
re.ReplaceAllStringFunc("ABCCCBBAAABC", func(s string) string { return replaceMap[s] })
```

### Write a Test

> [go.dev/add-a-test](https://go.dev/doc/tutorial/add-a-test)
> [gobyexample/testing](https://gobyexample.com/testing)

In short, to test a func `Fn` in package `mx.go`, create a file `mx_test.go` with same package, then add a `TestFn` function that takes in a `*testing.T` like so:

```go
func TestFn(t *testing.T) {
    want := "expected"
    got, err := mx.Fn();
    if err != nil {
        t.Fatalf(err) // <- Unexpected error
    }
    if want != got {
        t.Errorf("Expected ", want, ", got", got); // <-- Logical error
    }
    // Fatalf stops test run, Errorf shows error but continues other tests
    // Use each where it makes sense
}
```

You can then test this with `go test .` assuming a `go.mod` is in place.

### Compile time assertion for a type implementing an interface

TL;DR: To ensure that a type `T` satisfies interface `I`, add this one-liner:

```go
var _ I = (*T)(nil)
```

Long story -

Okay, this was a bit of a head-scratcher for me, so I'll explain it from the top as I understood it:

Premise: Go is "duck-typed" (like Python), in the sense you don't explicitly tell a type to implement an interface (like say, Java). Rather you write the methods as required by the interface and the type satisfies that interface implicitly.

```go
// Say we have an interface defined somewhere like so
type Xer interface  { X() }

// And a function that relies on the above
func DoX(x Xer) { x.X(); }

// In some other file, we have some type
type A struct {}

// And we want to pass an instance of it
// to the function that consumes the interface
func unrelatedCall() { a := A{}; DoX(a) }

// All it takes is for the type to implement an interface is to define the behaviour
func (a A) X() {}
```

Problem: This is all fine and dandy when the interface definition, consumer function and the struct itself are well-known to you (like in a small code-base). You can eyeball it to ensure everything is on the up and up. When you can't, how do you ensure and let others - humans and machines both, know as well that yes, `A` actually satisfies `Xer` and you can use it in functions that consume it?

Solution: That is where this magical line comes into play -

```go
var _ Xer = (*A)(nil)
```

Alright, let's break it down. Remember basic variable definition? `var <name> <type> = <value>` . Well, we don't actually need a `name` because we ain't gonna use it. So `_` it. `type` is the interface we're ensuring is satisfied by `value`. But we don't want to initialize a `value` either, so we make a `nil` pointer to it.

And thus, if `A` does not implement `Xer` by way of `X()` method, we will get an error. More importantly, any packages (and linters, lsps, etc) that use `A` now know that it implements `Xer`!

### Use `go run` as a shell script

> [stackoverflow.com/whats-the-appropriate-go-shebang-line](https://stackoverflow.com/questions/7707178/whats-the-appropriate-go-shebang-line)

Add this at the top of file (even before `package main`):

```bash
//usr/bin/env go run "$0" "$@"; exit "$?"
```

`chmod` and then your `./script.go` just works!

> <span style="color:orange">IMPORTANT!</span> go will complain if you have multiple files with `main` func

### Ignore multiple `main` functions in same package

### Structure a project?

### Use go:stringer, and when?

### Use go:embed, and when?

### Build a smaller binary?
