---
modified: Wed Sep  3 19:56:25 EDT 2025
---
# Typescript

> JavaScript with types

## 2023-01-27

- [Official Site](https://www.typescriptlang.org/)

- [Handbook](https://www.typescriptlang.org/docs/handbook/intro.html)

## Notes

- Default target is ES3, pretty old. Use `--target es2015` for browsers, and `--target esnext` for servers.

- Good to have `noImplicitAny` and `strictNullChecks` on.

- Always prefer using `interface` over `type` unless explicitly needed.

  - The former is extendable and mergable
  - It shows up by name in error messages.
  - `type` can alias existing types (or interfaces): `type Username = string | IUsername`.
  - `type ColorAndCircle = Color & Circle` is an intersection of two interfaces.

## Exhaustive Type Checking

Consider an enum:

```typescript
enum Color {
  Red,
  Blue,
  Green,
}
```

that has a switch case:

```typescript
function whichColor(c: Color): string {
  switch (c) {
    case Color.Red:
      return "Red";
    case Color.Blue:
      return "Blue";
    default:
      return "Unknown color";
  }
}
```

but we missed `Color.Green` and yet no compiler errors!

A simple solution (Typescript 4.9+):

```typescript
function whichColor(c: Color): string {
  switch (c) {
    case Color.Red:
      return "Red";
    case Color.Blue:
      return "Blue";
    default:
      c satisfies never;
      return "Unknown color";
  }
}
```

Which leads to:

<span style="color: red; font-family: monospace">Type 'Color.Green' does not satisfy the expected type 'never'.</span>
