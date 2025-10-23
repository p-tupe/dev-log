---
modified: Wed Nov 23 17:59:51 EST 2022
---
# Java

- [Official Site](https://www.java.com/en/)

- [Official Docs Java 19](https://docs.oracle.com/en/java/javase/19/)

- [Official Docs Java 18](https://docs.oracle.com/en/java/javase/18/)

- [Official Tutorial](https://docs.oracle.com/javase/tutorial/)

- [Dev.java tutorial/docs beginner](https://dev.java/learn/)

- [Central Connecticut State University: Introduction to Computer Science using Java](https://chortle.ccsu.edu/Java5/index.html)

- [Java Programming for Beginners – Full Course: youtube video from freecodecamp.org](https://www.youtube.com/watch?v=A74TOX803D0)

## Notes

### Data types

> 8 primitive types in total

- Integer numbers (1, 2, -5)

  - byte 1B
  - short 2B
  - int 4B
  - long 8B

- Real numbers (1.2, 0.53)

  - float 4B
  - double 8B

- char ('a', 'x') 4B

- boolean (true, false) 1B (only 1 bit is actually used)

> Below types are not primitive (are stored on heap instead of stack)

- String, Arrays, Classes, Objects, Interfaces ...

### Variables

- Syntax: `<modifier> <type> <name> = <expression>;`

- Example: `int x = 1;`

- Rules:

  - Starts with a letter or underscore
  - Does not contain punctuations
  - Is not a keyword
  - Is not more than 42 chars

- Conventions:

  - Use camelCase for variables and functions
  - Use TitleCase for classes and interfaces
  - Use ALL_CAPS for constants

- use modifier `final` for constants

### Operators

- Arithmetic `+` `-` `/` `*` `%`

- Logical `&&` `||` `!`

- Bitwise `<<` `>>` `&` `~` = complement `|` = inclusive OR `^` = exclusive OR

- Comparison `==` `<` `>` `<=` `>=` `!=`

- Other Operators `+` `++` `-` `--` `=` and so on.

> some operators (like `+`) are overloaded depending on the operads and position

### Printing

- Normal string output to screen: `System.out.print[ln]("...")`

- Formatted string: `System.out.printf("%-,a.bx\n", args, args);`

  - % = format specifier
  - \- = left align
  - , = Thousands separator for numbers
  - a = spaces before
  - .b = spaces after
  - x = s(String)/d(Int)/f(Float)/c(Char)

## User Input (Scanner Class)

```java
import java.util.Scanner;

Scanner ip = new Scanner(System.in);

System.out.print("Enter something: ");
String something = ip.nextLine();
System.out.println("Something is " + something);

// <type> someVar = ip.next[Int | Float | Double | ...]();
// For next word use "next" (space terminates)
// For next line use "nextLine" (newline/enter terminates)
// Careful when using nextLine _after_ other next* as they leave newline in buffer

// VALIDATION
System.out.print("Enter an invalid int: ");
if (ip.hasNextInt()) {
  int x = ip.nextInt();
  System.out.println("it's a valid int: " + x);
} else {
  System.out.println("it's an invalid int: " + x);
}

input.close();
```

### Type casting gotchas

- integers => 7 / 3 = 2

- real => 7.0 / 3 = 2.333...

- Addition of precision results in implicit casting (like assigning an int value to a double var)

- Loss of precision results in compiler error (like assigning a double value to an int var)

  - Explicit casting is required to ignore errors: `int x = (int) 1.0;`

- real variable for int expression => double x = 7 / 3 => x = 2.0

- System.out.println(3 + 2 + 1 + "") => 6

- System.out.println("" + 3 + 2 + 1) => 321

### Conditionals (Branching | Selection)

- if else

  Syntax:

  ```
  if (<expression>) {}
  else if (<expression>) {}
  else {}
  ```

  Example:

  ```java
  int num = 9;

  if (num % 2 == 0) System.out.println("Number is even.");
  else System.out.println("Number is odd.");
  ```

- switch case

  Syntax:

  ```
  switch (<expression>) {
    case <value1>: <stmts>; [break];
    case <value2>: <stmts>; [break];
    case <value3>: <stmts>; [break];
    default: <stmt>;
  }
  ```

  Example:

  ```java
  int num = 8;

  switch (num % 2) {
    case 0: System.out.println("Num is even"); break;
    case 1: System.out.println("Num is odd"); break;
    default: System.out.println("Error!");
  }
  ```

  - `case` falls through

### Loops

- for

  Syntax:

  ```
  for (<initialize>; <condition>; <increment>) {
      <stmts>
  }

  // OR

  for (<type> <var_name> : <collection of type>) {
    <stmts>
  }
  ```

  Example:

  ```java
  int [] arr = {1, 2, 3, 4, 5, 6};

  for (int x = 0; x < 5; x++) {
    System.out.println(arr[x]);
  }

  for (int x: arr)  System.out.println(x);
  ```

- while (there is also "do-while")

  Syntax:

  ```
  while(<expression>) {}
  ```

  Example:

  ```java
  int x = 10;

  while(x > 5) {
    System.out.println(x);
    x--;
  }
  ```

- can use `break` and `continue` to manipulate loops.

### Strings

- [API Link](https://docs.oracle.com/en/java/javase/19/docs/api/java.base/java/lang/String.html)

- Strings are immutable character sequences saved on heap (not primitive values). Inside `java.lang.String` so automatically imported.

- Some useful String methods:

  - length()
  - charAt(int i)
  - indexOf(char c)
  - substring(int start, int end)
  - toUpperCase(), toLowerCase()
  - contains(String subStr)
  - endsWith(String subStr)
  - startsWith(String subStr)
  - matches(String regex)
  - replace(String old, String new)
  - split(String regex)
  - strip() [evolved from trim(), unicode aware]

- Note that while `String s = "abc"` is similar to `String s = new String("abc")`, it is not exactly the same because the former "literal" will first check in the String pool if `"abc"` already exists before creating a new one while the latter will create a new one irrespectively. The difference is explained by following code:

  ```java
  String s1 = "abc";
  String s2 = "abc";

  // s1 == s2 => true

  String s3 = new String("abc");
  String s4 = new String("abc");

  // s3 == s4 => false
  ```

### Arrays

- [API Link](https://docs.oracle.com/en/java/javase/19/docs/api/java.base/java/util/Arrays.html)

- Arrays are present in package `java.util.Arrays`

```java

// Syntax

// <type> [] <var_name> = new <type>[<capacity>];

// Example

int [] array_of_five_ints = new int[5];

String [] arrayOf10String = new String[10];

MyObject [] arrayOfCustomObj = new MyObject[3];

// Example to create with pre-filled data

char [] vowels = { 'a', 'e', 'i', 'o', 'u' }; // => new Char[5];

```

### Classes & Objects

```java
/*

Syntax:

<modifier> class <class_name> extends <another_class> implements <interfaces> {

    // fields
    <modifier> <type> <var_name>;

    // constructor
    <modifier> <class_name>(<type> <var>) {
      <var_name> = <var>;
    }

    // methods
    <modifier> <type> <fn_name>(<params>) {
        ...
    }
}

modifier = public, private, protected, static etc

constructor is optional and can be overloaded like methods

*/

// Example Public Bike Class

public class Bike {

  static int numWheels = 2;
  int fuelTank = 12; // Ltr

  public Bike(int fuelTankInLtr) {
    fuelTank = fuelTankInLtr;
  }
}

// Creating a new Bike Object

Bike b = new Bike(18);

```

### OOP Pillars

- Encapsulation - Packing state/behaviour (data/methods) together

  - Facilitates "data hiding"

- Inheritance - State/Behaviour can be passed around to descendants

- Polymorphism - Same name (object/method) means different thing depending on context

  - Overloading: Same name, different parameters, same class
  - Overriding: Same name, same parameters, different class

- Above 3 are pillars, OOP also consists keywords like

  - Abstraction: Define a skeleton (Abstract Class) or expose API (Interface) then implement later however.
    - Abstract Class - When classes (extends) are similar
    - Interface - When classes (implements) share some traits but are not similar
