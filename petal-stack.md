---
modified: Tue May 24 13:57:49 EDT 2022
---
# PETAL Stack

> Phoenix Elixir Tailwind Alpine Liveview

## Setup

[Youtube Link](https://www.youtube.com/watch?v=vZBHkvTAb2U)

1. Install Elixir using package-manager `yay -S elixir`
   a. Check: `elixir -v`
2. Setup hex archive manager `mix local.hex && mix local.rebar`
   a. Check: `mix hex`
3. Install Phoenix `mix archive.install hex phx_new`
4. Create new project `mix phx.new <project_name>`
5. Setup tailwind
   1. `$ cd assets && npm init -y`
   1. `assets $ npm i -D tailwindcss autoprefixer postcss postcss-import`
   1. `assets $ npx tailwindcss init -p`
   1. Add `postcss-import: {}` to postcss.config.js at the top
   1. Add `mode: 'jit'` and `content: ["../lib/*_web/**.*ex", "./js/**/*.js"]` to tailwindcss.config.js
   1. Remove phoenix.css import and add three tailwind imports in app.css
   1. Remove app.css import in app.js
   1. In dev.exs, add watcher for tailwind:
      ```
      npx: [
        "tailwindcss",
        "-i=css/app.css",
        "-o=../priv/static/assets/app.css"
        "--postcss",
        "--watch",
        cd: Path.expand("../assets", __DIR__)
      ]
      ```
6. Setup alpine
   1. `assets $ npm i alpinejs`
   1. in app.js:
   ```js
   import Alpine from "alpinejs";
   window.Alpine = Alpine;
   Alpine.start();
   ```
7. In mix.exs `aliases`

   1. add `"cmd --cd assets npm install"` to `setup` tail
   1. add `"cmd --cd assets npm run deploy"` to `assets.deploy` head

8. In assets/package.json, add script: `"deploy": "NODE_ENV=production tailwindcss --postcss --minify -i css/app.css -o ../priv/static/assets/app.css"`

## [Elixir](./elixir.md)

## [Tailwind](./tailwindcss.md)

## [Alpine](./alpine.md)

## Phoenix

> Phoenix is a web development framework written in Elixir which implements the server-side Model View Controller (MVC) pattern.

### Links

- [Official Site](https://hexdocs.pm/phoenix/overview.html)

## Notes

Base setup includes:

- Database Mapper (Ecto, Adapter)
- Telemetry
- Mailer (Swoosh)
- Http/websocket connection server (Cowboy)
- Testing framework (Floki)
- Tasks to generate stuff (like auth, schemas etc) automatically `mix help --search "phx"`

Regarding Ecto:

1. Setup DB - config/config.exs (and other imports if auto-gen) and lib/<project>/repo.ex. An example for "test" sqlite3 setup:

```elixir
# config/config.exs
config :test,
  ecto_repos: [Test.Repo]

config :test, Test.Repo,
  database: Path.expand("../test_dev.db", Path.dirname(__ENV__.file)),
```

```elixir
# test/repo.ex
defmodule Test.Repo do
  use Ecto.Repo,
    otp_app: :test,
    adapter: Ecto.Adapters.SQLite3
end

```

2. Setup Schema and Migrations

Schema is what maps data from db to elixir format, and contains changesets.

Migrations are used to create/alter tables such that they approach the state specified by schemas.

Regarding Assets:

Phoenix uses esbuild to pack css and js from `assets/` dir and served from `priv/static` dir. Static assets are _not_ copied over automatically!

Regarding Templates:

Be careful with variables in rendered html. If rendering from inside functions, make sure to use `@variable` instead of `assign.variable`
