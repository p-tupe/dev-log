---
modified: "Tue Apr  7 17:42:01 EDT 2026"
---

# Javascript

- [MDN Docs](https://developer.mozilla.org/en-US/docs/Web/javascript)

- [Comprehensive Tutorial](https://javascript.info/)

- [A good breakdown of function terminology](https://dev.to/aaron_powell/named-function-vs-variable-function-11m0)

- [Microsoft Developer Javascript Resources](https://developer.microsoft.com/en-us/javascript/)

- [javascript.info](https://javascript.info/intro)

## How To

### Make a dynamic chain of promises

```js
// To convert
const arr = [d1, d2, d3, ..., dn];
const fn = async (d) => { /* returns a promise */ };
// into
Promise.resolve().then(() => fn(d1)).then(() => fn(d2)).then(() => fn(d3))...then(() => fn(dn)).catch(error)
// use
arr.reduce((c, d) => c.then(() => fn(d)), Promise.resolve()).catch(error);
// OR
for (const d of arr) await fn(d);
```

### Convert string to num shorthand

```javascript
const numStr = "1.23";
const numVal = +numStr; // 1.23
```

### Get integer part of a fraction (like Math.floor)

```js
const fraction = 1.234;
const intPart = ~~fraction; // 1
const intPart2 = fraction << 0; // 1
```

> The `~` operator is 2's complement

### Get an input from stdin

```javascript
import * as readline from "node:readline/promises";
import { stdin as input, stdout as output } from "node:process";

const rl = readline.createInterface({ input, output });
const answer = await rl.question("What do you think of Node.js? ");
console.log(`Thank you for your valuable feedback: ${answer}`);
rl.close();
```
