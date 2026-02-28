---
modified: "Sun Jan 18 12:18:38 EST 2026"
---

# html

- [MDN Docs](https://developer.mozilla.org/en-US/docs/Learn/HTML)

- [freeCodeCamp youtube link](https://www.youtube.com/watch?v=a_iQb1lnAEQ)

## How to adjust image src for light/dark mode?

```html
<picture>
  <source media="(prefers-color-scheme: dark)" srcset="dark-mode-image.png" />
  <source media="(prefers-color-scheme: light)" srcset="light-mode-image.png" />
  <img alt="Fallback image description" src="default-image.png" />
</picture>
```
