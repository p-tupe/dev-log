---
modified: Wed Nov  8 20:03:05 EST 2023
---
# Backup Strategy

> Everything worth backing up follows 3-2-1 strategy : Atleast 3 copies, 2 local and 1 remote

I currently have following storage options:

- Local

  - Laptop
  - Phone
  - External Disk

- Remote

  - Mega.nz: 50 GB
  - Oracle VPS: 50 GB
  - Nextcloud: 100 GB
  - Storj: 150 GB
  - Github

## Strategies:

- Multi-Media (Photos, Videos, Music, Ebooks and general Documents)

  - Laptop <-> Phone
  - Nextcloud (auto-sync)
  - External Disk (once a month)

- Notes, (Encrypted) Keys, etc

  - Laptop <-> Phone
  - Github (immediate)
  - Nextcloud (auto-sync)
  - External Disk (once a month)

- Project Repositories

  - Laptop (generated)
  - Github (immediate)
  - External Disk (once a month)

- Deployed Apps (Config and DB)

  - VPS local backups
  - Daily automatic encrypted push to Storj
  - Laptop
  - External Disk

## Notes:

- I started using Mega in days of yore when they still gave out 50 GB for new users. Most of my data currently lives there although I'm in the process of migrating to Nextcloud.

- I only _very_ recently discovered Storj as AWS alternative. It has a pleasant UI and easy config. I will be using it for all my automatic backups.

- For VPS (generated), I actually store backups elsewhere in a local dir (away from active app data). At any point in time, atleast 7 days' backup exists in any location.

- I use systemd timers (in lieu of cron jobs) for automatic backup scripts. And I have a slack channel that receives only failure notifications.

## Tools:

- [Syncthing]() Laptop <-> Phone
- [Rsync]() Laptop <- VPS
- [Rclone]() VPS -> Storj/Mega
- [Mega.nz App]() Laptop -> Mega.nz
- [Nextcloud]() for Laptop -> Nextcloud
- [Git]() Laptop -> Github
- [Cable/ADB]() Laptop <-> Phone <-> External Disk
