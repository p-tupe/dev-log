---
modified: Thu Oct 23 12:48:34 EDT 2025
---
# SSH

- Connect to remote host

  ```shell
  $ ssh <username>@<hostname/ip> [-p <port>]
  ```

- Setup keys (no password login)
  - Generate keys on client

    ```shell
    $ ssh-keygen -t ed25519  -f ~/.ssh/<key-name> -C "<Some comment>"
    ```

    where,  
     t = type of key algo  
     f = file name of generated keys  
     C = comment regarding who and where of key usage

  - Ensure passphrase is entered, it is remembered later

  - Copy public key to remote server

    ```shell
    $ ssh-copy-id  -i ~/.ssh/<key-name>.pub <username>@<hostname>
    ```

    where,  
     i = Identity file to use

  - Turn off password authentication on remote server

    ```shell
    $ sudo sed -i \
      -e 's/#\?PasswordAuthentication yes/PasswordAuthentication no/' \
      -e 's/PubkeyAuthentication no/PubkeyAuthentication yes/' \
      /etc/ssh/sshd_config
    ```

  - May also disable `PermitRootLogin`

  - May change default port

  - Reload ssh daemon via systemctl `sudo systemctl reload sshd`

  - May change default ssh porto

  - Should use tool like fail2ban to reject unauthorized attempts

  - To disable login banner/info:

    ```shell
    $ sed -i 's/PrintLastLog yes/PrintLastLog no/' /etc/ssh/sshd_config
    $ touch /home/$USER/.hushlogin
    ```

- Add a known host to ssh config for easier connection (also used by scp & rsync), using ssh config

  ```shell
  $ touch ~/.ssh/config && chmod 600 ~/.ssh/config
  ```

  ```config
  Host server-name # server-name is pattern matched
    HostName <hostname/ip> # indentation optional but recommended
    User <username>
    Port <port>
    PreferredAuthentications publickey
    IdentityFile ~/.ssh/<private-key>
  ```

  ```shell
  $ ssh server-name
  ```

- Recommended - generate a new ssh key pair for every remote host (so even if stolen, cannot compromise others). Can add in config as:

  ```config
  Host name1
    ...
    Identity ~/.ssh/<key-name>

  Host name2
    ...
    Identity ~/.ssh/<key-name>
  ```

## Running local scripts on remote host

- [Source](https://sqlpey.com/shell/execute-local-scripts-remote-ssh/)

```bash
ssh user@server 'bash -s' < local_script.sh
```

```bash
ssh user@server 'bash -s' <<'EOF'
# ...commands
whoami
EOF
```
