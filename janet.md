---
modified: Wed Jun  1 21:19:07 EDT 2022
---
# Janet

> Janet is a functional (and imperative) cross-platform interpreted (and compiled) lisp-inspired language with batteries-included api that makes a good system scripting language, or a language to embed in other programs.

> The entire language (core library, interpreter, compiler, assembler, PEG) is less than 1MB.

## Resources

[Official Site](https://janet-lang.org)

## Notes

### _Very_ Basics

> Starting off with minimum required to do anything

Data Types:

> Comments start with a `#`

- `nil`
- booleans: `true`, `false`
- symbols (conventionally `small-letters` with hypens)
- keywords: `:symbols` starting with colon
- numbers: like javascript number `1` `-1.1` `1e10` (unlike js, `4r100` = 16: `rn` = base `n`)

Data Structures:

- strings: `"strings are in double quotes"` or in n\`s like lua's `[[ ... ]]`
- buffer: mutable `@`strings
- tuples: `[1 2 3]` list of values
- arrays: mutable `@`tuples
- structs: like elixir's `{:keyword "value"}` pairs (or js objects, or clisp plists)
- tables: mutable `@`structs

Bindings:

> anything with paren tuple `(...)` is "called" (conventionally called "forms")

- Variable definition: `([def | var] <name> <expr>)`

  - immutable `def` or mutable `var`
  - bind any value retured by `<expr>` to a symbol `<name>`
  - works as variable assignment but can also destructure (pattern match) on lists and dicts
  - use `(let [<name> <expr> <name2> <expr2>] ...)` for multiple `def` bindings sugar

- Variable reassignment: `(set <name> <expr>)`

  - only works on `var`
  - `name` must exists and be an initialized value or a declared datastructure

- Function definition: `([defn | fn] [<name>] <args> <body>)`

  - yes, `name` is optional in case of `fn`. functions form closure.
  - arguments are literal tuples `[...]` but can be destructured from `struct`s
  - use `&` for variadic arguments: `[a b c & rest]`
  - body is zero or more forms `(...)`
  - last form is returned implicitly
  - `defn` is short for binding `fn` to `<name>` using `def`

Conditionals:

- If conditional: `(if <expr> <true-form> [<false-form>])`

  - lazy eval (does not evaluate forms until needed)
  - if no `false-form` but `cond` is `false`, then returns `nil`

- Cond conditional: `(cond <expr> (body) <expr2> (body2) <expr3> (body3))`

  - Can be used instead of multiple if-else-if-else nesting

- Case conditional:

  - think `switch` statements elsewhere

- When conditional:

  - like `if` without `false-form`; returns `nil` on `false` implicitly

Loops:

> OPINION: Avoiding mutable loops and stick to functional paradigm

- While loop: `(while <cond> <body>)`

  - Usually works with mutable `vars` with side effects in body

- For loop: `(for <var> <start-expr> <end-expr> <body>)`

- loop loop: `(loop [<var> <verb> <values> <modifiers> <cond>] <body>)`

  - Ex1: `(loop [i :range [0 10] :when (even? i)] (print i))`

    - `i` is a var, `:range` is a verb, `:when` is a modifier

  - Ex2: `(loop [name :in names] (print name))`

    - `name` is a var, `:in` is a verb and `names` is a list

  - Ex2 can be written using `each` and `map` resp as:

    - `(each name names (print name))`

    - `(map print names)`

  - Ex3: `(loop [[k v] :pairs some_struct] (print k " -> " v))`

    - Above for looping across a dict

Arithmetic `(<op> <expr> <expr2>)` and Comparisons `(<cmp> <expr> <expr2>)` work as usual:

- Use `(not= <expr> <expr2>)` for checking unequality
- Use `(deep= <mutable> <mutable>)` for checking contents of mutable structures
