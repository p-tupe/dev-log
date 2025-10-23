---
modified: Sun Jan 29 19:34:45 EST 2023
---
# Progressive Web Apps

> Web Apps that work (almost) like native apps on all platforms

- [Mozilla Docs](https://developer.mozilla.org/en-US/docs/Web/Progressive_web_apps)

## Notes

- In addition to an "App Shell", requires an "App Manifest" and "Service Worker"

- App Shell is the minimum code needed to render the UI, that can be cached. It's a mix of HTML, CSS and JS files.

- App Manifest is a json file that tells the browser about PWA and how it should appear on device.

  - Usually named `manifest.json` in root directory and linked as `<link rel="manifest" href="manifest.json" />` in all pages.

  - [Manifest Properties](https://developer.mozilla.org/en-US/docs/Web/Manifest)

- Service Worker is a sort of Web Worker specific to PWAs. It is responsible for all the Fetching, Caching, Notifying etc.

  - Lifecycle: Installing, Error, Activated, Idle, Fetch/Message, Terminated
