---
modified: "Thu Nov 13 20:51:07 EST 2025"
---

# Shell (bash | zsh)

## Using a default value for a variable

```bash
VAL=${<var>:-"<default>"}
```

The above sets default if var is null or unset. Use only `:` for skipping null check.

## Reuse previous shell command

```bash
!!
```

Example: check git root and cd if proper

```bash
git rev-parse --show-toplevel
# /git/root
cd $(!!)
```

## Reinstall xcode-select @MACOS

```bash
sudo rm -rf $(xcode-select --print-path) && xcode-select --install
```

## Mount external ntfs hdd @MACOS

1. `brew cask install osxfuse`
2. Reboot
3. `brew install ntfs-3g`
4. Connect hdd
5. `mkdir ~/NTFS`
6. `diskutil list` (check MS drive num eg disk2s1)
7. `sudo umount /dev/disk2s1`
8. `sudo /usr/local/bin/ntfs-3g /dev/disk2s1 ~/NTFS -olocal -oallow_other`

## Mount an attached block storage (pre-formatted) to instance using fstab

1. Check currently attached volumes `lsblk` (ensure attached but not mounted)
2. Get UUID of volume `blkid`
3. Add entry in /etc/fstab `UUID=<from above> /mount/point format options`
4. Mount all fstab volumes `sudo mount -a`
5. Check currently attached volumes `lsblk` (ensure mounted)

> <span style="color:orange">IMPORTANT!</span> If attaching new block storage, may need to format first

## Check size of a directory

```bash
du -sh <dir>
```

where,

- s = summary
- h = human readable format

Example: Get size of all directories in current directory

```bash
du -sh ./*/
```

## Move new files only (do not overwrite if exists)

```bash
mv -vn <source> <target>
```

where,

- v = verbose
- n = no overwrite

## if/else

```bash
if [ <test> ]; then <something>; fi
```

for command sucess | fail, simply do this (no sq brackets)

```bash
if command; then <sucess>; else <fail>; fi
```

To check if an environment variable exists

```bash
if [[ -z $ENV_VAR ]]; then echo "Not exists"; fi;
```

## for loop

```bash
for ((expr1;expr2;expr3)); do <command>; done;
```

Example: Loop from 1 to END

```bash
END=5
for ((i=1;i<=END;i++)); do echo $i; done;
```

## for in loop

```bash
for <variable> in <list>; do <command>; done
```

Example: Add all dot files to gitignore

```bash
for file in .*; do echo $file >> .gitignore; done
```

## Arrays

An array in shell is within `()` brackets and looped like so:

```bash
arr=("a" "b" "c")

for alph in ${arr[*]}; do echo $alph; done;
```

Note that on cli, direct `$arr` works for looping but not in script!

## Script strict mode

_Ignore the \ in \# below, kept for formatting reasons_

```bash
#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'
```

[http://redsymbol.net/articles/unofficial-bash-strict-mode/](http://redsymbol.net/articles/unofficial-bash-strict-mode/)
## Find command

```bash
find <path> -name "regex" [-exec <another cmd> | -delete]
```

Example: find and delete all dot files in current directory

```bash
find . -name ".*" -delete
```

as an alternative, `fd` works as follows

```bash
fd -t [d|f] [-HIL] "regex"
```

## Remove duplicate lines from a file

```bash
awk '!l[$0]++' file > newfile
```

if order is not important

```bash
sort -u file > newfile # u = unique
```

## Start a background job

```bash
nohup <job> < /dev/null > /dev/null 2>&1 & disown
```

where,

- nohup = no hangup signal
- \< /dev/null = no stdin (input)
- \> /dev/null = no nohup file (output)
- 2>&1 = redirect errors to stdout (which is null)
- & = start in background
- disown = detach from shell

Example: To start [mpv](https://mpv.io/) as music daemon

```bash
nohup mpv --no-audio-display --shuffle ~/Music < /dev/null > /dev/null 2>&1 & disown
```

## Mount Ram as disk

```bash
sudo mount -t tmpfs -o size=5g tmpfs /mnt/ramfs
```

where a new temporary mount point /mnt/ramfs is created with 5 Gigabytes of capacity;
to unmount, do

```bash
umount /mnt/ramfs
```

## Watch a file for change and run command when changes

```bash
while true; do watch -g ls -l <filename> && <cmd>; sleep 5; done
```

## Download a webpage and all it's linked pages and content

```bash
wget -r -l 1 -p -k -H -D domain.com,relateddomain.com http://domain.com/page/in/domain
```

where,

- r = recurse
- k = patch local links
- H = traverse domains other than original
- D = limit domains traversed
- l = depth of recursion
- p = download related content like images

## Generate a random base64 string (May be used as token or password)

```bash
head -n1 /dev/urandom | base64
```

- number after `n` changes length of string

## Get all lines _except_ the first:

```bash
tail -n+2
```

## Lazy regex quatifier (works in js, not is sed, vim)

Suppose there is a line like so:

`"this is some string" "this is another string"`

And the aims is to select everything within quotes as a string.

A naive regex would be: `/".*"/`

But it actually selects everything inside starting from first " to the end of line "

To select only the content withing first closing quote, do : `/.*?/`

That is, add a `?` aka a lazy quatifier (as opposed to normal greedy approach)

## Add colour to console output (make sure to add _Reset_ at the end after adding colour)

`\033[0m` : Reset
`\033[1m` : Bold
`\033[91m` : Red
`\033[92m` : Green
`\033[93m` : Blue

- JS Function

```js
const colorize = (...args) => ({
  black: `\x1b[30m${args.join(" ")}`,
  red: `\x1b[31m${args.join(" ")}`,
  green: `\x1b[32m${args.join(" ")}`,
  yellow: `\x1b[33m${args.join(" ")}`,
  blue: `\x1b[34m${args.join(" ")}`,
  magenta: `\x1b[35m${args.join(" ")}`,
  cyan: `\x1b[36m${args.join(" ")}`,
  white: `\x1b[37m${args.join(" ")}`,
  bgBlack: `\x1b[40m${args.join(" ")}\x1b[0m`,
  bgRed: `\x1b[41m${args.join(" ")}\x1b[0m`,
  bgGreen: `\x1b[42m${args.join(" ")}\x1b[0m`,
  bgYellow: `\x1b[43m${args.join(" ")}\x1b[0m`,
  bgBlue: `\x1b[44m${args.join(" ")}\x1b[0m`,
  bgMagenta: `\x1b[45m${args.join(" ")}\x1b[0m`,
  bgCyan: `\x1b[46m${args.join(" ")}\x1b[0m`,
  bgWhite: `\x1b[47m${args.join(" ")}\x1b[0m`,
});
```

## Process SIGnals (`man signal`)

- The INT (interrupt) signal is the weakest. It is the signal number 2. Its meaning is "stop what you're doing right now and wait for further user input". It's the signal generated by Ctrl+C in a terminal

- The TERM (terminate) signal is the normal kill signal (i.e. the one sent when you just call kill). It is signal number 15. It tells the application to exit cleanly. This signal allows for a trap handler, which enables the receiving process to do some clean up in an orderly fashion. An application that doesn't want to be interrupted during a critical operation might ignore SIGTERM for a while.

- The HUP signal (hang up) is about the same as SIGTERM in terms of harshness, but it has a specific role because it's automatically sent to applications running in a terminal when the user disconnects from that terminal. It is signal number 1.

- The QUIT signal is the harshest of the ignorable signals. It's meant to be used when an application is misbehaving and needs to be killed now, and by default it traditionally left a core dump file. It is signal number 3.

- The KILL signal does not have ANY possibility of the being resisted. It always works because a KILL signal cannot be handled. This is both its blessing and its curse. It is signal number 9.

## Using args

```bash
# Usage: argsFn 1 2 3 4 5 6 7 8
argsfn() {
  echo "number of args: $#"
  echo "list of args:"; echo $@
  echo "string of args: $*"
  echo "name of script: $0"

  if [ "$#" -lt 3 ]; then
    echo "usage: $0 arg1 arg2 arg3"
    exit
  fi

  echo "first arg: $1"
  echo "second arg: $2"
  echo "third arg: $3"

  shift 3
  echo "rest of the args: $*"
}
```

## Using flags [getopt](https://mywiki.wooledge.org/BashFAQ/035#getopts)

## Get just a filename from whole path

```bash
basename -s .txt some/path/to/file.txt
```

## SSHFS (macos)

```bash
brew install --cask macfuse
brew tap gromgit/fuse && brew install gromgit/fuse/sshfs-mac
sshfs -V # test if installed
sshfs -o allow_other,default_permissions <server-dir> <local-dir> # follow prompts for permission
```

> On linux, sshfs should be preinstalled, if not just use default package manager

## Notifications

```bash
# Usage: notify "Hello, World!"
notify() {
  if [[ `uname` == Darwin ]]; then
    osascript -e 'display notification "'$1'"'
  else
    # TODO linux
    # notify "Notification Text"
  fi
}
```
