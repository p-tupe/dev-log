---
modified: Thu Oct 23 12:48:34 EDT 2025
---
# Javascript

- [MDN Docs](https://developer.mozilla.org/en-US/docs/Web/javascript)

- [Comprehensive Tutorial](https://javascript.info/)

- [A good breakdown of function terminology](https://dev.to/aaron_powell/named-function-vs-variable-function-11m0)

- [Microsoft Developer Javascript Resources](https://developer.microsoft.com/en-us/javascript/)

## Notes

- Make a dynamic chain of promises

  Use

  ```js
  arr.reduce((c, d) => c.then(() => fn(d)), Promise.resolve()).catch(error);
  ```

  to convert

  ```js
  const arr = [d1, d2, d3, ..., dn];
  const fn = async (d) => { /* returns a promise */ };
  ```

  into

  ```js
  Promise.resolve().then(() => fn(d1)).then(() => fn(d2)).then(() => fn(d3))...then(() => fn(dn)).catch(error)
  ```

  - In newer versions:

    ```js
    for (const d of arr) await fn(d);
    ```

- Convert string to num shorthand

  ```javascript
  const numStr = "1.23";
  const numVal = +numStr;
  ```

- Get integer part of a fraction (like Math.floor)

  ```js
  const fraction = 1.234;
  const intPart = ~~fraction; // 1
  const intPart2 = fraction << 0; // 1
  ```

  > The `~` operator is 2's complement
