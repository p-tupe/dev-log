---
modified: "Fri Mar 20 10:28:36 EDT 2026"
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
// Word scanner version
var c string
fmt.Print("Continue? [y/n]: ")
fmt.Scan(&c)

// Line scanner version
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

### Write a Benchmark

> [b-loop](https://go.dev/blog/testing-b-loop)

In `whatever_test.go`, simply write:

```go
func BenchmarkX(b *testing.B) {
    for b.Loop() {
        // ... stuff to test
    }
}
```

then just `go test -bench=. whatever_test.go` to see the output.

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

```go
// Let's say we want B to implement X, but forgot to do so
type B struct{}
var _ Xer = (*B)(nil) // <-- this line will now error out!
```

### Use `go run` as a shell script

> [stackoverflow.com/whats-the-appropriate-go-shebang-line](https://stackoverflow.com/questions/7707178/whats-the-appropriate-go-shebang-line)

Add this at the top of file (even before `package main`):

```bash
//usr/bin/env go run "$0" "$@"; exit "$?"
```

`chmod` and then your `./script.go` just works!

> <span style="color:orange">IMPORTANT!</span> go will complain if you have multiple files with `main` func

### Ignore multiple `main` functions in same package

Best option I found is to simply keep packages in different directories as different modules; You can then `go install .` them to run from wherever.

### Structure a project

For all projects, start with:

```bash
mkdir project-name && cd project-name
go mod init whatever/project-name
touch project-name.go # package main for apps
```

If it's a short script, this is enough. If you need to split code, say separate some utilities:

```bash
mkdir internal/utils/utils.go # contains a function X
```

and inside `project-name.go`, import utils like so:

```go
// Import the internal utils package
import "whatever/project-name/internal/utils"
// And use the function like so:
utils.X()
```

If it's an CLI/API, by convention that code goes into a `cmd/cli` dir, while for a website, it goes into a `cmd/web` dir (sibling to `internal`).

### Gracefully shut down a go app

> https://go.dev/wiki/SignalHandling

```go
package main

func main() {
    // Make a context tha returns a cancel fn
	ctx, cancelJobs := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancelJobs()

    // Create a chan that wait on a signal
    sig := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)

    // At the end of main, wait on the chan for the signal
    <-sig
	log.Println("Shutting down gracefully")
}
```

### How to run an external command?

> See [go.dev/os/exec](https://pkg.go.dev/os/exec) for package docs

```go
// For normal run, but show error if not:
if err := exec.Command("echo", "hello, world!").Run(); err != nil {
    log.Println(err)
}

// For when combined (stdour/err) output is needed:
cmd := exec.Command("echo", "hello, world!")
if op, err := cmd.CombinedOutput(); err != nil {
    log.Fatalf("Error while running cmd: %v", err)
} else {
    log.Println(string(op))
}

// For when a cmd struct is necessary (for customization):
cmd := exec.Cmd{
    // See https://pkg.go.dev/os/exec#Cmd
    Path:   "/bin/echo",
    Args:   []string{"echo", "hello,", "world"},
    Stdout: os.Stdout,
}
if err := cmd.Run(); err != nil {
    log.Fatalf("Error while running cmd: %v", err)
}
```

### Accept an array of options in go flag

```go
// This is a simpler way with no parsing
flag.Parse()
inputs := flag.Args() // inputs contains all options
```

- [Source](https://beta.stackoverflow.com/questions/28322997/how-to-get-a-list-of-values-into-a-flag-in-golang)

```go
// This is a more involved one
package main

import "flag"

type arrayFlags []string

// String is an implementation of the flag.Value interface
func (i *arrayFlags) String() string {
    return fmt.Sprintf("%v", *i)
}

// Set is an implementation of the flag.Value interface
func (i *arrayFlags) Set(value string) error {
    *i = append(*i, value)
    return nil
}

func main() {
    var myFlags arrayFlags
    flag.Var(&myFlags, "list1", "Some description for this param.")
    flag.Parse()
}
```

### Run a fn every X seconds

```go
ticker := time.Tick(3 * time.Second)
for range ticker { /*...*/ }
```

## Slices gotchas

- Ranging on a slices creates a _copy_ of each item; so modifying a slices goes like so:

```go
for idx, item := range allItems {
    // This won't work, will only modify a locally copied item
    item.key = value

    // Use this instead to modify the slice
    allItems[i].key = value
}
```

- Variadic args do not pass a copy, don't modify assuming a clone

```go
func x() {
    og := []int{1, 2, 3}
    y(og...)
    fmt.Println(og) // 1, 2, 9
}

func y(arr ...int) {
    arr[2] = 9 // arr points to same underlying memory as og
    fmt.Println(arr)
}
```

- Slices delete function _zeroes elements_ in place and _returns_ a modified slice.

```go
sliceA := []int{2, 4, 3, 1, 6}
// Slice A here:  [2 4 3 1 6]

sliceB := slices.DeleteFunc(sliceA, func(i int) bool {
    return i == 4 || i == 1
})
// Slice A after: [2 3 6 0 0]
// Slice B after: [2 3 6]

sliceC := slices.Delete(sliceB, 1, 2)
// Slice A now: [2 6 0 0 0]
// Slice B now: [2 6 0]
// Slice C now: [2 6]
```
