---
modified: Sun Jan 29 19:34:45 EST 2023
---
# Kotlin

- [Official site](https://kotlinlang.org/)

- [Android Dev Docs](https://developer.android.com/kotlin/learn)

## Notes

- `main` function is the primary entry point

- Print to standard output: `print("Hello, World!")`

- Datatypes (almost same as Java)

  - `Int` `Long` `Short` `Byte`
  - `Double` `Float`
  - `String` `Char`
  - `Boolean`

- Variables and Constants

  - Constants: `val <name> [: <type>] = <value>`
  - Variables: `var <name> [: <type>] = <value>`

- Operators

  - Logical: `&&` `||`
  - Bitwise: `and` `or`
  - Equality: `===` `!==`

- Data Structures

  - List: `val lst = listOf("a", "b", "c")` List<String>
  - Map: `val mp = mapOf("a" to 1, "b" to 2, "c" to 3)` Map<String, Int>
  - Set: `val st = setOf("a", "b", "c")` Set<String>
  - Add `mutable` as prefix for variables: `val mutList = mutableListOf("a")`

- Functions

  ```kotlin
  // Syntax
  // fun <name>([... <param>: <type>]): <type> { ... }
  // fun <name>([... <param>: <type>]) = ...
  
  // Example 1
  fun sum(a: Int, b: Int) = a + b
  
  // Example 2
  // Unit means void, can be omitted
  fun sum(a: Int, b: Int): Unit {
      print("$a and $b")
  }
  ```

- Loops

- Branching

  - `when` is like switch:

    ```kotlin
    when (x) {
      1 -> print("1")
      2 -> print("2")
      else -> print("None")
    }
    ```

- Optional

  - Use `?` for nullable type: `var x: String? = null`
  - Use `?` in call chain for safe calls: `print(x?.length)`
  - Use `!!` to throw if null (like force unwrap `x!` in Swift): `print(x!!.length)`
  - Use `?:` to get a default value (like `x ?? <default>` in Swift): `print(x?.length ?: 0)`
  - Use `as?` for safe casting: `val anInt: Int? = x as? Int`
  -

- Class/Objects
