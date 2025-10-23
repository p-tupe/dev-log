---
modified: Wed Jun 19 14:08:30 EDT 2024
---
# Docker

- [Archwiki Doc](https://wiki.archlinux.org/title/Docker)

- [Docker Docs](https://docs.docker.com/get-started/)

- [Docker container resource management](https://dzone.com/articles/docker-container-resource-management-cpu-ram-and-i)

- A "container" is a virtually isolated environment. An "image" is formed from multiple layer of commands.

- Analogy : An "image" is like the concept of Class from OOP - a blueprint/recipe of what will be. A "container" on the other hand, is like an Object - an instance of Class, or the dish made from a recipe. Obviously, there can be multiple "containers" for same "image".

- Start docker service

  ```sh
  systemctl start docker.service
  ```

- Start now & enable automatic startup on login

  ```sh
  systemctl enable --now docker.service
  ```

- Seek help

  ```sh
  docker [command] help
  ```

  ```sh
  docker <command> [subcommand] --help
  ```

- Deploy a container

  ```sh
  docker run <image-name> [<command>]
  ```

  with following options:

  - -d = detached mode
  - -e = environment variables
  - -p <host port>:<container port> = port redirect
  - -v <host dir> | <volumne name>:<container dir> = bound volume | named volume
  - -i = interactive mode (keeps STDIN open)
  - -t = pseudo-tty
  - -name = assign a name to the container
  - -w = work directory inside container
  - -rm = automatically remove on exit

  to get the rest of the options possible, run

  ```sh
  docker run --help
  ```

- Show deployed containers

  ```sh
  docker ps [-a for all]
  ```

- Execute a command in deployed container

  ```sh
  docker exec [-it for interactive tty] <container-id> <command>
  ```

- Stop a container

  ```sh
  docker stop <container-id>
  ```

- Remove a container

  ```sh
  docker rm <container-id>
  ```

- Force stop & remove a container

  ```sh
  docker rm -f <container-id>
  ```

- Pull an image

  ```sh
  docker pull <image-name>
  ```

- Show all images

  ```sh
  docker images
  ```

- Remove an image

  ```sh
  docker rm <image-name> [or <image-id> for unnamed images]
  ```

- Build an image

  ```sh
  docker build -t [<image-namespace>/]<new-image-name>[:<image-tag>] .
  ```

  assuming current directory contains Dockerfile.

- Create a named volume

  ```sh
  docker volume create <some-name>
  ```

- List all volumes

  ```sh
  docker volume ls
  ```

- Use host's network interface (reduces NAT latency)

  ```sh
  docker run ... --net=host ...
  ```

- When mounting volumes "${PWD}" works, "${pwd}" doesn't. Keep env var case sensitivity in mind.

- Can get instance properties of containers (eg ip) as by runnings inspect

  ```sh
  docker inspect my-container
  ```

- Add current user to docker group (to avoid typing 'sudo' on every docker command); recommended only on trusted machines.

  ```sh
  sudo gpasswd -a $USER docker
  ```

- To access localhost (outside container), use ip address of bridge interface

  ```sh
  ip addr show docker0
  ```

- To Change default logger and data-root directory, add following to /etc/docker/daemon.json

  ```json
  {
    "log-driver": "json-file",
    "log-opts": {
      "max-size": "10m",
      "max-file": "3"
    },
    "data-root": "/mnt/volume/docker"
  }
  ```

  - Note that local log driver does not show logs with docker compose
  - Other notable options instead of local: journald, syslog,

### Troubleshoot

- Permission denied error solution: `sudo aa-remove-unknown`

> It turned out that AppArmor service was messing up with Docker. AppArmor (or "Application Armor") is a Linux kernel security module that allows the system administrator to restrict programs' capabilities with per-program profiles.

### MacOS

> Dockers require a linux kernel to work, so it works like a virtual machine on Mac.

- Easiest way to use would be `brew install --cask docker` but docker desktop is not opensource (or fully free).

### Deploying static sites in docker

Following dockerfile uses nginx:alpine image to serve static data over 8000

```dockerfile

# nginx state for serving content
FROM nginx:1.21-alpine

# Set working directory to nginx asset directory
WORKDIR /usr/share/nginx/html

# Remove default nginx static assets
RUN rm -rf ./*

# Copy static assets over
COPY ./static ./

```

Build the image:

```sh
docker build -t static-serve .
```

Start a container

```sh
docker run -p 8000:80 static-serve
```

---

Btw, this is all an alternative to:

```sh
python -m http.server
```

---

Also see multistage builds to reduce image size, if that is a concern. See [how-to-use-the-official-nginx-docker-image](https://www.docker.com/blog/how-to-use-the-official-nginx-docker-image/)

---

For smallest possible image size for static serve, see [smallest-docker-image-static-website](https://lipanski.com/posts/smallest-docker-image-static-website)

## Copy Docker image without hub/registry

```sh
docker save <image> | gzip | DOCKER_HOST=ssh://user@remotehost docker load
```
