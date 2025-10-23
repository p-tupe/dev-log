---
modified: Tue May 24 13:57:49 EDT 2022
---
# Browser APIs

## Tips n Tricks

```js
location.href = "/some/url"; // <-- This simulates a link click, allowing back navigation
location.replace("/some/url"); // <-- Does not do the above
```

## Web Sockets

```js
// Connected to backend web socket server
const ws = new WebSocket("ws://whatever:port/and/path");

// Receive message from server
ws.onmessage = (event) => console.log("Received from websocket: ", event.data);

// Send message to server
ws.send("whatever");
```

## Web Workers

From main script:

```js
// Start a new Web Worker (`service-worker.js`)
const worker = new Worker("path/to/service-worker.js");

// Received messages from Worker
worker.onmessage = (event) => console.log("Main: ", event.data);

// Send message to Worker
worker.postMessage(["Hello", "from", "main!"]);
```

From worker script:

```js
onmessage = function (e) {
  console.log("Worker: ", e.data.join(" "));
  postMessage("hello from worker!");
};
```

### Shared Web Workers

- Shared by any script using same protocol, host and port

- Mostly similar to dedicated Worker, except communication happens via "port" explicitly (instead of assumed in case of dedicated)

From main script:

```js
// Start a new Web Worker (`shared-worker.js`)
const worker = new SharedWorker("path/to/shared-worker.js");

// Received messages from Worker
worker.onmessage = (event) => console.log("Main: ", event.data);

// Send message to Worker
worker.postMessage(["Hello", "from", "main!"]);
```

From shared-worker script:

```js
onconnect = function (e) {
  let port = e.ports[0];
  port.onmessage = function (e) {
    console.log("Worker: ", e.data.join(" "));
    port.postMessage("Hello from worker!");
  };
};
```

### Broadcast Channels
