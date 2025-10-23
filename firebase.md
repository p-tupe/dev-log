---
modified: Tue Feb 28 05:03:18 EST 2023
---
# Firebase

- [Official Docs](https://firebase.google.com/docs/database)

- [SQL vs Firebase DB](https://www.youtube.com/watch?v=WacqhiI-g_o)

## Realtime DB (For Web)

- Is a single JSON object

- Use Firebase SDK (platform-specific) to interface with database: `const db = firebase.database().ref()`

- Get a reference to a node (key): `const ref = db.child("path/to/ref")`

- To retrieve data in real-time, add event listeners like so:

  ```js
  // value events work best for syncing objects
  objectRef.on("value", snap => {...})

  // child events work best for lists
  listRef.on("child_added", snap => {...})

  // Use ref.off to stop listening
  const handle = ref.on("value", ...)
  objectRef.off("value", handle)
  ```

  - EVENTS

    - value
    - once
    - child_added
    - child_removed
    - child_changed
    - child_moved

- To query data, the syntax is usually `reference().order().query1().query2()...on()`, like so:

  ```js
  ref.orderByChild("someKey").equalTo("someValue").limitToFirst(1).on(...)
  ```

  - ORDER

    - orderByKey() // eg paging
    - orderByChild() // eg where clause in SQL
    - orderByValue() // eg numeric sorting

  - QUERY

    - startAt() // range start
    - endAt() // range end, use with startAt for unicode support (see eg below)
    - equalTo()
    - limitToFirst()
    - limitToLast()
    - startAfter()
    - endBefore()

  ```js
  // Example node
  // "users": {
  //     "1": {
  //         "name": "abc",
  //         "age": 123
  //       }, ...
  //   }

  // Get all "users" that have "name" starting with "X"
  ref.child("users").orderByChild("name").startAt("X").endAt("X\uf8ff");

  // Get all users whose age is less than 100
  ref.child("users").orderByChild("age").endAt(99);
  ```

- Multipath (ATOMIC) Update

  ```js
  // for updateObject like
  const updateObject = {
    "path/to/ref/1/key": "newValue",
    "path/to/another/ref/key": "newValue",
    "path/to/one/more/ref/key": "newValue",
  };

  // Update all values (starting from root) at once (or fail together)
  rootRef.update(updateObject);
  ```

- Security Rules

  - A JSON file separate from DB

    ```js
    {
      "rules": {
        "some/ref": {
          "$childKey": {
            ".read": true,
            ".write": true,
            ".validate": true
          }
        }
      }
    }
    ```

    - Use server variables to construct rules:

      - Auth
      - Data
      - New Data (inside .write and .validate)
      - Now (timestamp)
      - Root
      - $wildcards

      ```js
      // Users can only see/modify their own data, where their name must be a string
      {
        "rules": {
          "users": {
            "$uid": {
              ".read": "auth.uid == $uid",
              ".write": "auth.uid == $uid"
              ".validate": "newData.child('name').isString()"
            }
          }
        }
      }
      ```
