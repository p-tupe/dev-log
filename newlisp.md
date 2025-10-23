---
modified: Sat Feb 18 13:36:03 EST 2023
---
# newLISP

- [Official Site]()
- [Introduction to newLISP]()
- [newLISP in 21 minutes]()

## Notes

- List: `(x y z)` -> S-expression

  - Example: `(1 2 3 4 5)` -> List of whole numbers until 5

- Evaluation: `(f x y z)` -> First symbol is evaluated (considered to be a function)

  - Example: `(+ 1 2 3 4 5)` -> Addition on list of whole numbers until 5

- Quotes: `'(1 2 3 4 5)` -> Do not evaluate first symbol, returns a list as is

  - Example: `(println '(+ 1 2 3 4 5))`

- Global Symbols: Like variables/functions; References to S-expressions

  - Example: `(define x 6)` | `(define l '(1 2 3 4 5))`

- Comments: Comments start with `;` until EOL

  - Example: `(define n 42) ; the meaning of life`

- Branching:

  - if:

    - Syntax: `(if <cond> <true> [<false>] [... <cond> <true> [<false>]])`
    - Example: `(if should-reroute? (switch-to-other))`
    - Example: `(if (> 4 2) "4 > 2" "4 < 2") ; "4 > 2"`
    - Example: `(if (> 2 1) "2 > 1" (> 3 2) "3 > 2" "3 < 2") ; "2 > 1"`
    - `nil` | `()` | unknown are false, rest true

  - when: Like an `if` without the `else`, can execute multiple exprs if true

    ```lisp
    (when (> 1 0)
      (println "First body")
      (println "Second body")
      "I'm returned")
    ```

  - case: Switch to an expression based on value of case

    ```lisp
    (case v
      (1 (print "ONE"))
      (2 (print "TWO"))
      (3 (print "THREE"))
      (true (print v)))
    ```

    - Case values are literals, not expressions; they are not evaluated

  - cond: check if conditions are true, then evaluate (mix of if and when)

    ```lisp
    ; Syntax: (cond (<cond> <exprs>) [... (<cond> <exprs>)])
    (cond
      ((> 1 x) (println "First") (println "Second"))
      ((> 10 x) (println "First Again") (println "Second Again"))
      (true (println "Catch-all")))
    ```

- Looping:

  - dolist: Loop on a list

    ```lisp
    ; Syntax: (dolist <list> <exprs>)
    (dolist (i (sequence 0 10))
      (println "Element " i " at index " $idx))
    ```

  - dostring: Loop on a string

    ```lisp
    ; Syntax: (dolist <string> <exprs>)
    (dostring (code "abcdef") (println code))
    ```

  - dotimes: Loop for a count

    ```lisp
    ; Syntax: (dotimes <counter> <exprs>)
    (dotimes (c 10) (println c)) ; (c from 0 to <count>), here 0..9
    ```

  - for: Loop for a "custom" count

  ```lisp
  ; Syntax: (for <custom counter> <exprs>)
  (for (c 5 -5 .5) (println c)) ; (c <from> <to> <step>)
  (for (c 1 10) (println c)) ; here 1..10
  ```

      - break out early: Use a test condition to exit loop

        ```lisp
        ; Syntax: (looper (<counter> <test>) <exprs>) ; Stop if test is true
        ; Works in dolist, dostring and for as well
        (dotimes (c 10 (= c 5)) (println c)) ; Stop at count 5
        ```

  - while (do-while): Loop for a condition success (opposite of until false)

    ```lisp
    ; Syntax; (while <success-cond> <exprs>)
    (while (disk-not-full?) (add-file))
    ```

  - until (do-until): Loop for a condition failure (opposite of while true)

    ```lisp
    ; Syntax: (until <failure-cond> <exprs>)
    (until (disk-full) (add-file)) ; Does the same as above while
    ```

  - map: Apply a function over a list

    ```lisp
    ; Syntax: (map <fn> <list>)
    (map println (sequence 0 10))
    ```

- Lambas:

  ```lisp
  (fn (i) i * i) ; An anonymous function that returns square of i
  ```

- Blocks: Groups of related but independent S-expressions; can be groups explicitly

  - begin: Group expressions into a single list, return value of last

    ```lisp
    (begin (first) (second) (third) (fourth) (return-fifth-value))
    ```

  - and: Finish all expressions and return truthy or fail one and return nil immediately

    ```lisp
    (and (first) (second) (third-fails-returns-null) (fourth-doesnt-run))
    ```

  - or: Finish expressions until a truthy value; return nil if none of them return true

    ```lisp
    (or (first-gives-nil) (second-gives-nil) (third-is-true) (fourth-not-evaled))
    ```

  - amb: chose a random (ambiguous) expressions to eval

    ```lisp
    (amb 1 2 3 4 5) ; returns a new number every time
    ```

- Local Symbols: Variables valid only inside the S-expression they are defined in

  ```lisp
  ; (let (<sym> <expr>) <exprs using sym>)
  (let (x 2 y 3 z 4) (println x " " y " " z))

  ; letn (<sym1> <exprs> <sym2> <exprs using sym1>) <exprs using sym>)
  ; Used when symbols are used during initialization
  ```

- Functions: Using "define"

  ```lisp
  ; Syntax: (define (<name> <args>) <exprs>)
  (define print-something (println "Something"))
  ```

  - Should use `let` for local variables; use `set` to change values of local variables

  - `args` contains list of undeclared arguments to a function; `doargs` can be used to iterate through them
