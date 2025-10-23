---
modified: Wed Jul 16 11:16:20 EDT 2025
---
# systemd

- [Official Site](https://systemd.io/)

> Systemd is a suite of basic building blocks for a Linux system. It provides a system and service manager that runs as PID 1 and starts the rest of the system.

- [A blog from creator](https://0pointer.de/blog/projects/systemd.html)

- [Arch Wiki](https://wiki.archlinux.org/title/Systemd)

- [Systemd by example](https://systemd-by-example.com/)

- [Red Hat Docs](https://docs.redhat.com/en/documentation/red_hat_enterprise_linux/9/html/using_systemd_unit_files_to_customize_and_optimize_your_system/assembly_working-with-systemd-unit-files_working-with-systemd)

## Units

- Services (Daemons/Units) [Uses `systemctl`]
- Timers (similar to Cron)
- Mounts (File system mounts, from /etc/fstab)
- Devices
- Sockets

- List all units for a type:

  ```sh
  systemctl list-units --type <name>
  ```

## Serices

- Start a service:

  ```sh
  systemctl start <name>
  ```

- Stop a service:

  ```sh
  systemctl stop <name>
  ```

- Enable a service to start on boot:

  ```sh
  systemctl enable <name>
  ```

- Disable a service from starting on boot:

  ```sh
  systemctl disable <name>
  ```

- Disable a service from being started by other services:

  ```sh
  systemctl mask <name>
  ```

- Status of a service:

  ```sh
  systemctl status <name>
  ```

  - Enable [or disable or mask] and start immediatly:

  ```sh
  systemctl enable [disable | mask] --now <name>
  ```

  - Reload service units

  ```sh
  systemctl daemon-reload
  ```

## Journal

- [Arch Wiki](https://wiki.archlinux.org/title/Systemd/Journal)

> systemd has its own logging system called the journal; running a separate logging daemon is not required. To read the log, use `man journalctl(1)`

- Show all messages since boot:

  ```sh
  journalctl -b [-n for nth boot, 0 = current, 1 = previous etc]
  ```

- Show all messages with prority higher than n [2 = critical, 3 = error, 4 = warn]:

  ```sh
  journalctl -p <n>
  ```

- Show all messages for a unit:

  ```sh
  journalctl -u <unit-name>
  ```

  - If unit is in `~/.config/` then instead of `-u`, use `--user-unit`

- Show all messages for an executable:

  ```sh
  journalctl /path/to/executable
  ```

- Clear up old logs

  ```ssh-config
  # Check size
  journalctl --disk-usage

  # Rotate logs
  sudo journalctl --rotate

  # Remove all logs beyond 7d
  sudo journalctl --vacuum-time=7d
  ```

## Timers

- [Arch Wiki](https://wiki.archlinux.org/title/Systemd/Timers)

> Timers are systemd unit files with a suffix of .timer. Timers are like other unit configuration files and are loaded from the same paths but include a [Timer] section which defines when and how the timer activates.

> For each .timer file, a matching .service file exists (e.g. foo.timer and foo.service). The .timer file activates and controls the .service file.

### Example of a timer service:

- [Website link #Demo](https://fedoramagazine.org/systemd-timers-for-scheduling-tasks/)

- Executable to run (can be any language, using a bash script here). Inside $HOME/bin/schedule-test.sh:

  ```sh
  #!/usr/bin/env bash
  echo "This is only a test: $(date)" >> "$HOME/schedule-test-output.txt"
  ```

  - Make sure the executable is... executable.

    ```sh
    chmod u+x $HOME/bin/schedule-test.sh
    ```

- Systemd Service Unit inside ~/.config/systemd/user/schedule-test.service:

  ```ini
  [Unit]
  Description=A job to test the systemd scheduler

  [Service]
  Type=simple
  ExecStart=/home/<user>/bin/schedule-test.sh

  [Install]
  WantedBy=default.target
  ```

- Systemd Timer Unit inside ~/.config/systemd/user/schedule-test.timer:

  ```ini
  [Unit]
  Description=Schedule a message every 1 minute

  [Timer]

  #Execute job if it missed a run due to machine being off
  Persistent=true

  #Run 120 seconds after boot for the first time
  OnBootSec=120

  #Run every 1 minute thereafter
  OnUnitActiveSec=60

  #File describing job to execute
  Unit=schedule-test.service

  [Install]

  # So systemd can manage automatically
  WantedBy=timers.target
  ```

  - Can use `OnCalendar` instead of `OnUnitActiveSec`

- Reload daemons and start service to test:

  ```sh
  systemctl --user daemon-reload
  systemctl --user start schedule-test.service
  tail $HOME/schedule-test-output.txt
  ```

- Enable timer if test pass and check status:

  ```sh
  systemctl --user enable --now schedule-test.timer
  systemctl --user status schedule-test.timer
  tail -f $HOME/schedule-test-output.txt
  ```

- If want to run sudo, put files in `/etc/systemd/system` and use `sudo systemctl`

- If want to run as user but need sudo as well, add `User=username` in [Service]

## Systemd service order

```
local-fs-pre.target
         |
         v
(various mounts and   (various swap   (various cryptsetup
 fsck services...)     devices...)        devices...)       (various low-level   (various low-level
         |                  |                  |             services: udevd,     API VFS mounts:
         v                  v                  v             tmpfiles, random     mqueue, configfs,
  local-fs.target      swap.target     cryptsetup.target    seed, sysctl, ...)      debugfs, ...)
         |                  |                  |                    |                    |
         \__________________|_________________ | ___________________|____________________/
                                              \|/
                                               v
                                        sysinit.target
                                               |
          ____________________________________/|\________________________________________
         /                  |                  |                    |                    \
         |                  |                  |                    |                    |
         v                  v                  |                    v                    v
     (various           (various               |                (various          rescue.service
    timers...)          paths...)              |               sockets...)               |
         |                  |                  |                    |                    v
         v                  v                  |                    v              rescue.target
   timers.target      paths.target             |             sockets.target
         |                  |                  |                    |
         v                  \_________________ | ___________________/
                                              \|/
                                               v
                                         basic.target
                                               |
          ____________________________________/|                                 emergency.service
         /                  |                  |                                         |
         |                  |                  |                  To do this, we first need to add an [Install]
         v                  v                  v                                 emergency.target
     display-        (various system    (various system
 manager.service         services           services)
         |             required for            |
         |            graphical UIs)           v
         |                  |           multi-user.target
         |                  |                  |
         \_________________ | _________________/
                           \|/
                            v
                  graphical.target
```
