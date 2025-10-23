---
modified: Mon Jul 18 18:32:52 EDT 2022
---
# NUXT v3

NOTE: v3 is currently a RC, link may change when stable
[Official Site](https://v3.nuxtjs.org/)

## SETUP

### Automatic

```shell
$ npx nuxi init <new-project>
```

### Manual

Create a new project

```shell
$ mkdir <new-project> && cd $_
$ npm init -y
$ npm install nuxt@rc # this install nuxt 3 (release candidate)
$ touch app.vue nuxt.config.js tsconfig.js
```

Add the following to package.json

```json
  "scripts": {
    "dev": "nuxt dev",
    "build": "nuxt build",
    "generate": "nuxt generate",
    "start": "nuxt start"
  }
```

## NOTES

`app.vue` is the "starting-point" of nuxt app. Technically, we can already build a nuxt project as long as `app.vue` and `nuxt.config.js` is present
