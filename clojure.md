---
modified: Wed Nov 23 17:59:51 EST 2022
---
# Clojure(script)

## Resources

- [Clojure Official Site](https://clojure.org/index)
- [Clojurescript Official Site](https://clojurescript.org/index)

- [Clojure Resources #1](https://clojuredocs.org/)
- [Clojure Resources #2](https://clojure-doc.org/)

- [Clojure for the Brave and True:](https://www.braveclojure.com/)

- [Learn Clojurescript](https://www.learn-clojurescript.com/section-0/)

- [Modern CLJS](https://github.com/magomimmo/modern-cljs)

- [Clojurescript Unraveled](https://funcool.github.io/clojurescript-unraveled/)

- [Clojure from the ground up](https://aphyr.com/posts/301-clojure-from-the-ground-up-welcome)

- [Clojure by example](https://kimh.github.io/clojure-by-example/#about-this-page)

- [Clojurescript Koans](http://clojurescriptkoans.com/)

- [Exercism Track: Guided Tutorials](https://exercism.org/tracks/clojure)

- [Clojure Beginner Resources: Github Gist](https://gist.github.com/yogthos/be323be0361c589570a6da4ccc85f58f)

## Notes (from braveclojure)

### Syntax

> (operator operand1 operand2 ... operandn)

NOTE: The above expression is called a "form" (it is common to all Lisps)

OPINION: Should just call 'em expressions and be done with it

### Data Types

- `nil`
- Booleans: `true`, `false`
- Numbers: integers `1`, floats `1.2`, ratios `1/2`
- Strings: `"only double quotes work"`
- Characters: `\a` `\b` `\c`
- Regex: `#"regular-expression-here"`
- Keywords: `:colon-prefixed-symbols`
  - Very fast lookup, so used as keys and params
- Null: `nil`
- Symbols: `+` `map`

### Data Structures

- Maps (Dicts/Objects/Hashes): `{:keyword-keys "value can be anything"}`

  - Get: `(get {:key "value"} :key "default")` => Lookup by key, "default" if not found
  - Add: `(assoc some-map :key "V")`
  - Update: `(assoc some-map :key "U")`
  - Remove: `(dissoc some-map :key)`
  - Contains: `(contains? some-map :key)` => true | false
  - Find: `(find some-map :key)` => [:key "U"]
  - Keys: `(keys some-map)`
  - Values: `(vals some-map)`
  - Merge: `(merge map-one map-two)`
  - Get-in: `(get-in {:a {:b "value"}} [:a :b])` => Nested
  - `(:key {:key "value"} "default")` == `({:key "value"} :key)` == `"value"`
  - Assoc-in updates nested data

- Records (Alternatives to maps for domain specific data with "type"):

  - Definition => `(defrecord Name [field1 field2 field3])`
  - Instance => `(def inst1 (->Name val1 val2 val3))`
  - Instance => `(def inst2 (map->Name {:field1 val1 :field2 val2 :field3 val3}))`

- Vectors (Arrays): `[1 2 3 "anything" :goes]`

  - Get: `(get [1 2 3] 0)` => Lookup by index
  - Add: `(conj [1 2 3] 4)` => Adds to **end** of vector

- Lists (Linked Lists): `'(1 2 3 "anything" :goes)`

  - Get head: `(first '(1 2 3))` => 1
  - Get tail: `(rest '(1 2 3))` => 2 3
  - Get top: `(pop '(1 2 3))` => 2 3
  - Peek top: `(peek '(1 2 3))` => 1
  - Get arbitrary: `(nth '(1 2 3) 0)` => Lookup by iterating
  - Add: `(conj '(1 2 3) 4)` => Adds to **start** of list

- Sets: `#{"some" :values 1 1/4}`

  - Add: `(conj #{:a :b} :c)` => :a :b :c
  - Remove: `(disj #{:a :b :c} :c)` => :a :b
  - Contains: `(contains? #{1 2 3} 2)` => Check inclusion, returns true/false
  - Can use Map-like `get` and `:key` lookup too

### Operators

> `(are_forms? x y z)` => `true`

- Equality: `(= x y)`
- Logical:
  - `(or x y)` => First truthy or first false/nil
  - `(and x y)` => First false/nil or last truthy
- Scope: `(do x y z)`

### Bindings

> `def` is all where all bindings start

- `(def const "value")` => symbol `const` is now bound to `"value"`

- `(defn fun-name "document string" [param & rest] (print "This is some form") {:last-form "returned"}) `

  - Functions can be "arity overloaded"
  - Function parameters can be "destructured"

- `(fn [arg] (print arg))` == `#(print %)` ==> anon fn

  - `fn` has closure over surrounding lexical scope

- `(let [name value] (code that uses name))` => `let` creates bindings in lexical scope

### Branching

- If-else: `(if <cond> <then> <else>)` => all things inside `<>` are individual forms

  - If no `else` then returns `nil` on `false`
  - Use `do` to add more than one expression in the branches

- When: `(when <cond> <form1> <form2> <form3>)`

  - A combination of `if` and `do`

- Cond: `(cond <condition1> <expression1> <condition2> <expression2>)` => Successively check each condition and evaluate corresponding expression if true

- Cond-else: `(cond <> <> <> <> :else <expression>)` => `:else` always evaluates to true. Using the word "else" is a convention.

- Case: `(case <var> <case_1_literal> <expression> <case_2> <expression>)`

  - Matches are done in constant time, case conditions must be literal values

- Case-else: `(case <var> <> <> <> <> <default expression if no match>)`

### Loops

- For: `(for <[bindings]> <expression>)`

  - More of a list comprehension than a loop. Returns a list.

- Loop + Recur: `(loop <[bindings]> (<expr> (recur <[bindings]>)))`

- Fn + Recur: `(defn name <[bindings]> (<expr> (recur <[bindings]>)))`

  - `recur` must be the last expression
  - it is tail call recursive

### Exceptions

### With (Same as `with` in python)
