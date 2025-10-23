---
modified: Mon Jul 18 18:32:52 EDT 2022
---
# rclone

* [Official Site](https://rclone.org/)

> Rclone is a command line program to manage files on cloud storage.

* I followed the steps for [Storj](https://storj.com) [here](https://docs.storj.io/dcs/how-tos/sync-files-with-rclone/rclone-with-hosted-gateway/)
* To add/edit a cloud provider

  ```
  rclone config
  
  ```

  and follow the prompts

## Bucket commands

* Create a bucket

  ```
  rclone mkdir <provider>:<new-bucker-name>
  
  ```
* List all buckets

  ```
  rclone lsf <provider>:
  
  ```
* Delete an empty bucket

  ```
  rclone rmdir <provider>:<bucket-name>
  
  ```
* Delete non-empty bucket

  ```
  rclone purge <provider>:<bucket-name>
  
  ```

## Object commands

* Copy an object from SRC to DST

  ```
  rclone copy [--progress] SRC DST
  
  ```
  * SRC/DST can be local/remote.
* List objects

  ```
  rclone ls <provider>:<bucker>/path
  
  ```
  * use lsf for non-recursive
* Delete an object

  ```
  rclone delete[file] <provider>:<bucket>/path
  
  ```
* Sync SRC and DST

  > Use the sync command to sync the source to the destination, changing the destination only. Doesn’t transfer unchanged files, testing by size and modification time or MD5SUM. Destination is updated to match source, including deleting files if necessary.

  ```
  rclone sync [--progress] SRC DST
  
  ```

## Password inside shell script

```
RCLONE_CONFIG_PASS=whateverthepasswordis
```

```
export $(cat /home/ubuntu/nextcloud/.env | xargs -0)
# ... rclone commands
unset RCLONE_CONFIG_PASS
```