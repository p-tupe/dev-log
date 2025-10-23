---
modified: Wed Jun 19 14:08:30 EDT 2024
---
# elixir

## Resources

- [Official Site](https://elixir-lang.org/)

- https://www.erlang.org/doc/system/getting_started.html

- https://hexdocs.pm/elixir/introduction.html

- https://hexdocs.pm/phoenix/overview.html

- [Official Learning Resource](https://github.com/dwyl/learn-elixir)

- [Standard Modules](https://hexdocs.pm/elixir/Kernel.html)

- [Elixir School: Online guide/tutorial](https://elixirschool.com/en/why)

- [Learn Elixir With Me: Blog Series](https://inquisitivedeveloper.com/tag/lwm-elixir/)

- [Exercism Track](https://exercism.org/tracks/elixir)

- [Koans](https://github.com/elixirkoans/elixir-koans)

- [Programming Elixir - Dave Thomas: Book](https://doc.lagout.org/programmation/Functional%20Programming/Programming%20Elixir_%20Functional%2C%20Concurrent%2C%20Pragmatic%2C%20Fun%20%5BThomas%202014-10-19%5D.pdf)

- [Learn you an erlang for greater good: Some concepts apply to elixir](https://learnyousomeerlang.com/content)

- [Reddit Community Link](https://www.reddit.com/r/elixir/)

- https://joyofelixir.com/toc.html

- https://www.poeticoding.com/category/elixir/

## TDD

Using ExUnit: `mix test --stale --max-failures 1 --listen-on-stdin --trace --seed 0`

- --stale executes only tests for source and test files which are changed.
- --max-failures 1 stops execution after first failing test.
- --listen-on-stdin allows you press Enter and have tests re-run.
- --trace prints names of tests instead of dots, so you can see which tests are executed in more descriptive way.
- --seed 0 forces tests to execute in same order they are mentioned in source file, so that you will have predictable consecutive runs of the tests.

### Recommended deps

- :mix_test_watch, Re-run test on file change, use options from above
- :ex_unit_notifier, Send test result notification, Combine with above

## Notes

### Data Types

- String
  - Interpolation: "#{var} literal"
  - Concatenation: "first" <> var <> "third"
- Integer
- Float
- Boolean
- Atom (like a "Symbol" in js, value is itself)

- Bitstring (just sequence of bits) > Binary (Bitstring % 8 === 0) > Strings (UTF-8 Binary)

### Data Structures

Tuples: `{"any", "type", 1, false}`

- Contiguous Memory Structure
- Cheap Read
- Expensive Write

List: `["here", "too", 123, true]`

- Linked List
- Cheap Write
- Expensive Read
- Append: [1 | [2, 3, 4]] = [1, 2, 3, 4]
- Concat: [1, 2] ++ [2, 3] = [1, 2, 3, 4]
- Substract: [3, 4] -- [1, 2, 3, 4] = [1, 2]

Charlist: `['a', 'b', 'c'] = 'abc'`

- List of Integers representing valid Unicode code points

Maps: `%{"any" => :value, 123 => true, "keys" => :too}`

- Cheap addition/deletion
- Expensive Read
- When key is atom, shortcut: `%{a: 1, b: 2} = %{:a => 1, :b => 2}` and `map.a = 1`
- Read key:
  - `map[key]`; `nil` on invalid key; useful for dynamically created maps
  - `map.key` when key is atom; error on invalid key; useful for static maps
- Update map:
  - Update existing `%{map | existing_key: new_value }`
  - Add new value `Map.put(map, key, value)`

Keywords: ?

Structs: ?

### Control Flow

- if: `if condition do \n ... \n end` or `if condition, do: ..., else: ...`

- unless: opposite of `if`, same syntax

- case: `case variable do \n pattern -> return \n ... \n end`

- cond: `cond do \n condition -> return \n ... \n end`

- with: `with pattern do \n ... \n end` or `with pattern, do: ...`

- Boolean Operators: and/&& or/|| not/!

- Exists: var in list, e.g, `1 in [1, 2, 3, 4]` = true, can be used as `not in`

### Functions

- Anonymous: `sum = fn a, b -> a + b end` called as `sum.(1, 2)` = 3
- Shorthand: `sum = &(&1 + &2)` same as above
- Parameters can be pattern matched to return different results:

  - `sum -> fn \n pattern -> return \n pattern -> return \n end`

### Modules

> Like Classes

- Functions inside modules are defined using `def` keyword, `defp` for private

- constants: `@name value`

- (parameters) structs: `defstruct [:var1, :var2]`

### Typespec and Doc

Create new type and assign it to spec:

```elixir
@type address_map() :: %{street: String.t(), postal_code: String.t(), city: String.t()}
@type address_tuple() :: {street :: String.t(), postal_code :: String.t(), city :: String.t()}
@type address() :: address_map() | address_tuple()

@doc """
Formats the address as an uppercase multiline string.
"""
@spec format_address(address()) :: String.t()
def format_address(%{street: street, postal_code: postal_code, city: city}) do
  format_address({street, postal_code, city})
end

def format_address({street, postal_code, city}) do
  """
  #{String.upcase(street)}
  #{String.upcase(postal_code)} #{String.upcase(city)}
  """
end
```

List of types: [Typespecs](https://hexdocs.pm/elixir/typespecs.html)
