---
modified: Wed Mar  2 19:04:17 EST 2022
---
# Nginx

- [Beginner's Guide](https://nginx.org/en/docs/beginners_guide.html#control)

- Some CLI commands:

  - Start nginx `sudo nginx` (takes config from /etc/nginx/nginx.conf)
  - Manage service : `sudo nginx -s [reload|stop|quit]`
  - Test config: `sudo nginx -t`
  - Custom config: `sudo nginx -c /custom/config` # Couldn't get this to work

- Config file syntax:

  - Consists of "directives" that are like for eg a "simple" directive: `<command> [<key>] <value>;`
  - A "block" directive consits of a number of simple directives in braces for eg:

    ```
    <block_command> <value> {
      <command1> <value>;
      <command2> <value>;
      <command3> <value>;
    }
    ```

  - Braces form a "context", which is similar to a lexical scope of javascript braces
  - And just like javascript global scope, nginx's is called "main context".
  - main { events {} http { server { location {} }}}

- Static Serve:

  - Define inside several server (and location) blocks
  - client request -> parse URI -> match longest location string -> append <value> and respond
  - Eg, 2 locations : /data/html /data/images with following config:

    ```
    server {
      location / {
        root /data/html;
      }

      location /images {
        root /data;
      }
    }
    ```

  request : localhost/ => matches / => /data/html/index.html
  request : localhost/images/test.png => matches /images (longest) => /data/images/test.png

  - Regular expression:

    ```
    location ~ \.(gif|jpg|png)$ {
      root /data/images;
    }
    ```

- Proxy:

  - Act as a middleman for requests
  - Client request -> Forward to specific server -> Receive reply from server -> Client response
  - Eg, request come on /proxy-this and need to be sent to a separate server on port 8080, and /proxy-that on port 8081, then config would be:

    ```
    server {
        listen 80;
        server_name somedomain.com; # This is optional
        location /proxy-this/ {
          proxy_pass        http://127.0.0.1:8080;
        }
        location /proxy-that/ {
          proxy_pass        http://127.0.0.1:8081;
        }
    }
    ```

  - Note that the trailing slash in proxy_pass value changes where location is truncated

  - Helpful directives:

    ```
    proxy_http_version 1.1;
    proxy_set_header Upgrade $http_upgrade;
    proxy_set_header Connection 'upgrade';
    proxy_set_header Host $host;
    proxy_set_header  X-Real-IP        $remote_addr;
    proxy_set_header  X-Forwarded-For  $proxy_add_x_forwarded_for;
    proxy_cache_bypass $http_upgrade;
    ```

- Load balancing: This config will start balancing requests on 3 addresses (default round robin)

  ```
  http {
    upstream poxy-server {
      server server_addr1;
      server server_addr2;
      server server_addr3;
    }

    server {
      location / {
        proxy_pass http://proxy-server;
      }
    }
  }
  ```
