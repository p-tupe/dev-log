---
modified: Wed Jul 16 11:16:20 EDT 2025
---
# Caddy

> Caddy is a web server with automatic certificate renewal and http to https routing.

## Resources

- [Official Link](https://caddyserver.com/)

- [Docker Image](https://registry.hub.docker.com/_/caddy)

## Setup

> I use it with Caddyfile and docker-compose.yml for easy config and deploy of caddy as a reverse-proxy (like [nginx](./nginx.md)).

This sets up the basic framework:

```shell
$ mkdir caddy
$ cd caddy
$ mkdir site
$ touch Caddyfile docker-compose.yml
$ docker create volume caddy_data
$ docker create volume caddy_config
```

[Docker](./docker.md) compose be like:

```yml
version: "3"

services:
  caddy:
    image: caddy:2
    network_mode: host
    volumes:
      - $PWD/Caddyfile:/etc/caddy/Caddyfile
      - $PWD/site:/srv:ro
      - caddy_data:/data
      - caddy_config:/config
    restart: unless-stopped

volumes:
  caddy_data:
    external: true
  caddy_config:
    external: true
```

[Caddyfile](https://caddyserver.com/docs/caddyfile-tutorial) for reverse-proxy is:

```nginx
# The entirety of ngnix nextcloud config turned into this:

nextcloud.mydomain.com {
  rewrite /.well-known/carddav /remote.php/dav
  rewrite /.well-known/caldav /remote.php/dav
  reverse_proxy localhost:8080
}
```

Finally, to run things:

```shell
$ docker-compose up -d
```

Done!

## Tip n Tricks

Reloading Caddyfile without downtime:

```shell
$ docker exec -w /etc/caddy caddy_caddy_1 caddy reload
```

Adding multiple reverse-proxy routes on same domain:

```nginx
api.mydomain.com {
    route /routeone/* {
        reverse_proxy localhost:8080
    }

    route /routetwo/* {
        reverse_proxy localhost:8081
    }
}
```

Adding basic authentication:

```shell
$ docker run -it --rm -w /etc/caddy caddy:<ver> caddy hash-password
```

This gives us a hased password, which we add like so:

```nginx
secureroute.mydomain.com {
  basicauth * {
      username <hash-string>
  }
}
```
