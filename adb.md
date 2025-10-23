---
modified: Wed Mar  2 19:04:17 EST 2022
---
# ADB

- Connect wirelessly

  1. Connect adb with cable
  2. `adb tcpip 5555`
  3. `adb connect <phone-ip>:5555`

- Push new files only

  ```sh
  adb push [-n for dryrun] --sync __source__ __target__
  ```

- Make an app fullscreen

  ```sh
  adb shell settings put global policy_control immersive.full=com.package
  ```

  where immersive.full can be replaced by immersive.status or immersive.navigation, to remove

  ```sh
  adb shell settings put global policy_control immersive.off=com.package
  ```

- Screenshot & Screenrecord

  ```sh
  adb shell screenrecord /sdcard/test.mp4
  ```

  Close with Ctrl-C

  ```sh
  adb shell screencap /sdcard/test.png
  ```

  To pull file on pc and remove from phone

  ```sh
  adb pull /sdcard/test.mp4 ./ && adb shell rm /sdcard/test.mp4
  ```
