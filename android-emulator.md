---
modified: Thu Oct 23 12:48:34 EDT 2025
---
## Android emulator cli

- [Source](https://developer.android.com/studio/run/emulator-commandline)

- Emulator bin path (macos): `$ANDROID_HOME/emulator/emulator`

- List available devices: `emulator -list-avds`

- Start an emulator: `emulator @<device_name>`

- Some useful options: `emulator @<device> -no-boot-anim -memory <size> -logcat <regex>`
