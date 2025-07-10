# Go

- [Official Site](https://golang.org)

- [Official Learning Resource](https://go.dev/learn/)

- [Go101: Beginners' Book](https://go101.org/)

- [Golang Tutorial: YouTube video tutorial](https://www.youtube.com/watch?v=YzLrWHZa-Kc)

- [How to Write Go Code](https://go.dev/doc/code)

- [Effective Go](https://go.dev/doc/effective_go)

- [Go By Example](https://gobyexample.com/)

- [Awesome Go](https://github.com/avelino/awesome-go)

- https://dev.to/aurelievache/learning-go-by-examples-introduction-448n

- https://golangbot.com/learn-golang-series/

- https://github.com/dariubs/GoBooks

- https://goperf.dev/

## Notes

### From "A tour of Go"

| module > package > functions

- Starting a go project:
  - Make a project dir: `mkdir project_dir`
  - Make a src dir: `mkdir package_name` (This will be the name of final executable)
  - Initialize go modules: `go mod init project_dir/package_name`

### From LinkedIn Learning "Learning Go" tutorial

- Go has 3 things (from OOP perspective):

  1. Type - Definition of a thing
  2. Struct - Data of a thing
  3. Interface - Behavior of a thing

  Everything in Go has a type, every type has an interface.

- Go _doesn't_ have these:

  1. No (type) inheritance (Also, no classes/objects, duh)
  2. No overloading
  3. No "structured" exception handling (try/catch/finally)
  4. No implicit casting

- Case-sensitive syntax; If first letter of symbol is Uppercase, then it is publicly available (exported)

- Builtin functions: [golang.org/pkg/builtin](https://golang.org/pkg/builtin)

- Types:

  - Everything has a type
  - Some built-in data types: `bool` `string` `[u]int(8 | 16 | 32 | 64)` `float(32 | 64)` `complex(64 | 128)`
  - Some built-in data type aliases: `[u]int` = `[u]int(32 | 64)` `byte` `uintptr` `rune`
  - Some built-in collection types: Arrays, Slices, Maps, Structs
  - Some built-in language types: Functions, Interfaces, Channels, Pointers
  - Use printf's `%T` to find type of something
  - Note that strings are in double quotes `""` and _characters_ (int32 | byte) in single `''` like Java

- Variables:

  - Declaration: `var <name> <type>`; Example: `var someString string`
  - Initialization: `<name> = <value>`; Example: `someString = "this is a string"`
  - Shortcut: `<name> := <value>`; Example: `anotherStr = "this is another string"`
    - Shortcut only works inside functions.

- Constants:

  - Definition: `const <name> <type> = <value>`
  - Cannot be reassigned

- `new()` vs `make()`

  - `new` is used to return a pointer of a "type" to an uninitialized memory location
  - `make` is used specifically with slices, maps and chans to allocate and initialize memory; it returns the object (as opposed to the pointer)

- Pointers:

  - Declare a pointer: `var ptr *int`
  - Declare & Initialize a pointer to val: `ptr := &val`
  - Address of val : `&val` == `ptr`
  - Value of val : `val` == `*ptr`

- Arrays:

  - Syntax: `var <name> = [<capacity>]<type>`, Example: `var rgb = [3]string`
  - Shorter syntax with Initialization: `rgb := [3]string{"red", "green", "blue"}`
  - Arrays are not usually preferred, since they cannot be sorted or capacity-changed

- Slices:

  - Slices used Arrays behind the scene, but are preferred over Arrays because their capacity can be changed at runtime
  - The syntax for slices is same as arrays, except the omission of `<capacity>` in square brackets
  - Another syntax: `<name> := make([]<type>, <init_size>, [<capacity>])`
  - If `capacity` is added, slice behaves like an array
  - Slice functions:
    - `append(<slice>, <value>)` - Add `<value>` to `<slice>`
    - `append(<slice>[<from>:<to>])` - Splice the slice
    - `sort.<type>(<slice>)` - Sort slice

- Maps:

  - Like Object in JS, Dict in Python, Hash Table
  - Syntax: `<name> := make(map[<key_type>]<value_type>)` (Here, `[]` does not mean optional, is part of the definition)
  - Example: `countryCodes := make(map[string]string)` `countryCodes["CA"] = "Canada"`
  - Map functions:
    - `delete(<map_name>, <key>)` - deletes `key` from `map_name`

- Structs:

  - Like a Class in Java, Can encapsulate both data and methods (optional)
  - Example:

    ```go
    type dog struct {
      age int
    }

    d := dog{5}
    fmt.Printf("%+v\n", d)
    d.age += 1
    fmt.Printf("%+v\n", d)
    ```

- Channels:

- If Statements:

  ```
  // Syntax (no brackets around conditions)
  if <condition> {}
  else if <condition> {}
  else {}
  ```

- Switch Statements:

  ```
  // Syntax
  switch <expression> {
      case <val>: <statements>
      ...
      default: <statement>
    }
  ```

  - case don't fall through by default (they break if satisfied)
  - must add `fallthrough` explicitly at the end of case statements

- For Loops:

  - Syntax 1 (same old):

    ```
    for i := 0; i < <some_int>; i++ { }
    ```

  - Syntax 2 (for _each_):

    ```
    for x, y := range <map | slice> {
      // x = key | index
      // y = value
    }
    ```

  - Syntax 3 (while like):

    ```
    for <condition> {
      // break and continue work as expected
    }
    ```

- Functions:

  ```
  func <name>(<args>) <return_type> {}
  ```

- I/O:

  - Console:

    ```go
    ip := bufio.NewReader(os.Stdin)
    fmt.Print("Enter something: ")
    name, _ := ip.ReadString('\n')
    fmt.Println("You entered: ", name)
    ```

  - File:

- Math:

  - `import "math"`
  - [Go Math Package](https://pkg.go.dev/math)
  - Arithmetic built-in: `+` `-` `*` `/` `%`
  - Bitwise built-in: `&` `|` `^` `&^` `<<` `>>`

- Date/Time:

  - `import "time"`
  - [Go Math Package](https://pkg.go.dev/time)
  - `time.Now()`
  - `time.Date(...)`

## Build an executable without debug (smaller size)
