---
modified: Sun Oct 30 19:36:58 EDT 2022
---
# Nodejs

## "main" module

There is a widly followed convention in python to have a "main" function that is run when the script is called directly via cli: `python script.py`

```py
def main():
  pass

if __name__ == "__main__":
  main()
```

A relatively unknown implementation exists for nodejs, like so:

```js
function main() {}

if (require.main === module) main();
```

<span style="color:red;font-weight:bold">IMPORTANT:</span> This only works for single run scripts and commonjs modules! ECMA Style "import" give module as undefined.

## Websocket

- low-level interface is net.Socket as an point for tcp connection

## Quick Server

```js
require("http")
  .createServer((req, res) => {
    console.log(req.method + ": " + req.url);
    res.end("OK");
  })
  .listen(8080);
```

## Testing

- A built-in testing framework using TAP format

  ```js
  import { describe, it, before, after } from "node:test";
  import { strict as assert } from "node:assert";

  describe("Test suite 1", () => {
    before("Before the tests", () => {});

    after("After the tests", () => {});

    it("should pass synchronous tests", (t) => {
      assert(true);
      // ...
    });

    it("should also pass async tests", async (t) => {
      // ...
    });
  });
  ```

- Use `node --test [<file_name>]` to run tests. Checks the followings (non-exhaustive list):
  - All `*.js` files in `test/` folder
  - All files starting with `test-*.js`
  - All files ending with `*-test.js` or `*.test.js`

## Readline

> The node:readline module provides an interface for reading data from a Readable stream (such as process.stdin or fileReader) one line at a time.

- To create a simple console user input prompt:

  ```js
  const readline = require("node:readline/promises");
  const { stdin: input, stdout: output } = require("node:process");

  async function main() {
    const rl = readline.createInterface({ input, output });
    const answer = await rl.question("Gimme some answer: ");
    console.log(`Thank you for your valuable feedback: ${answer}`);
    rl.close();
  }

  if (require.main === module) main();
  ```

- To create a simple file processor reading one line at a time:

  ```js
  const readline = require("node:readline/promises");
  const { createReadStream } = require("node:fs");

  async function main() {
    const input = createReadStream("input.txt");
    const rl = readline.createInterface({ input });

    for await (const line of rl) console.log(line);
    rl.close();
  }

  if (require.main === module) main();
  ```

- Although it can also be done like so:

  ```js
  import { open } from "node:fs/promises";

  const file = await open("./some/file/to/read");

  for await (const line of file.readLines()) {
    console.log(line);
  }
  ```

## About Timers

- [Offical event loop timers doc](https://nodejs.org/en/docs/guides/event-loop-timers-and-nexttick/)

- `process.nextTick(callback)` runs callback **immediately** after current synchronous code is done executing. It basically used to turn sync code into async so it follows proper flow. Use with caution as recursive nextTicks could starve the loop.

- `setImmediate(callback)` runs callback **after** the "polling" phase of the event loop, during a phase called "check".

- `setTimeout(callback)` and `setInterval(callback)` run callback during the "timer" phase of event loop.

- If both immediate and timeout (with 0 delay) are set in a synchronous context, either of them could run first. If they are set in async context (like inside a callback), then immediate runs first, since "check" phase comes after "poll" while "timer" phase will be first on the next loop.

- Also, nextTick is actually run literally immediately, during **any** phase of the loop, while setImmediate runs only during check phase. A historical misnomer.

## Worker Threads

## Measuring Performance
