---
modified: Wed Feb  9 17:51:07 EST 2022
---
# reactjs

- Use absolute imports `@component/Button`

  - As opposed to relative imports, which grow longer the deeper one goes
  - Makes it easier to refactor components
  - Makes code more readable
  - Will need more configuration for both bundling and ide errors

- Use builtin reducer to manage complex state `useReducer`

- Use index.js in components/ to import all common modules from one place instead of from each file

  - Helps avoid duplication
  - Reduces cognitive load
  - Easier to reason about code
  - Reduces loc

- Use "views" folder to hold single page components and "components" (or "modules") to define common basic components

- Wrap vendor components such that they are indistinguishable from custom components

- All styles, tests and data required by a view (or component) should be in one single folder

  - Makes it easier to view the whole project structure
  - Makes it trivial to add new components

- Split bundle by view (or route), load on-demand. Network requests are the slowest part of app life-cycle.
