---
modified: "Fri Nov 28 12:00:52 EST 2025"
---

# SSH

## Setup SSH keys (no password login)

1. Generate keys on client for remote-host

```shell
ssh-keygen -t ed25519  -f ~/.ssh/remote-host -C "Key for Remote Host"
```

> If you enter a passphrase, make sure to add key to ssh-agent for convenience

2. Copy public key to remote server

```shell
ssh-copy-id  -i ~/.ssh/remote-host.pub <user>@<host>
```

> This is just one of many ways to do it

3. That's all! Now `ssh user@remote-host` to login

## Tips

### Secure your server

Inside `/etc/ssh/sshd_config`, following changes are recommended:

```sshd_config
PasswordAuthentication no
PubkeyAuthentication yes
PermitRootLogin no
```

For extra security, change to a non-standard port for ssh (22 is the standard). If you have a public IPv6 address, use that instead of IPv4. If you can, change the username.

All this so scripted attacks (that scour wellknown usernames on standard ip/ports) can be mitigated.

Then reload ssh daemon via systemctl: `sudo systemctl daemon-reload && sudo systemctl restart sshd`

May also have something like `fail2ban` running; Ensure your firewall is up and working.

### Disable login banner/info

```shell
sed -i 's/PrintLastLog yes/PrintLastLog no/' /etc/ssh/sshd_config
touch /home/<user>/.hushlogin
```

### Quicker connection from client

Add a known host to ssh config for easier connection (also used by scp & rsync), using ssh config

```shell
touch ~/.ssh/config && chmod 600 ~/.ssh/config
```

```config
Host remote-host
  HostName <hostname/ip>
  User <username>
  Port <port>
  PreferredAuthentications publickey
  IdentityFile ~/.ssh/remote-host
```

```shell
ssh server-name
```

> It is recommended to generate a new ssh key pair for every remote host.

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
