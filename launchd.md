---
modified: "Mon Dec 29 17:29:14 EST 2025"
---

# launchd

> Linux systemd == Macos launchd

- Manuals: `man launchctl` | `man launchd`

- [An article on medium covering some basics](https://medium.com/swlh/how-to-use-launchd-to-run-services-in-macos-b972ed1e352)

- [Another article with more info](https://www.maketecheasier.com/use-launchd-run-scripts-on-schedule-macos/)

## Scheduling a script (zettelmerken) to run at specific time

1. Create a new file `com.zettelmerken.dailyreview.plist` in `~/Library/LaunchAgents/`

1. Add the following xml to it:

   ```xml
   <?xml version="1.0" encoding="utf-8"?>
   <!DOCTYPE plist PUBLIC "-//Apple Computer//DTD PLIST 1.0//EN"
   "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
   <plist version="1.0">
   <dict>

     <key>Label</key>
     <string>com.zettelmerken.dailyreview</string>

     <key>ServiceDescription</key>
     <string>Zettelmerken Daily Review</string>

     <key>ProgramArguments</key>
     <array>
       <string>/opt/homebrew/bin/python3</string>
       <string>-m</string>
       <string>zettelmerken</string>
     </array>

     <!-- Run on load (at bootup) -->
     <key>RunAtLoad</key>
     <true />

     <!-- Run Daily at 00:10 min -->
     <key>StartCalendarInterval</key>
     <dict>
       <key>Hour</key>
       <integer>0</integer>
       <key>Minute</key>
       <integer>10</integer>
     </dict>

     <!-- For Debugging
     <key>StandardErrorPath</key>
     <string>/tmp/com.zettelmerken.dailyreview.err</string>
     <key>StandardOutPath</key>
     <string>/tmp/com.zettelmerken.dailyreview.out</string>
     -->
   </dict>
   </plist>
   ```

1. Verify `plutil ~/Library/LaunchAgents/com.zettelmerken.dailyreview.plist`

1. Load agent `launchctl enable ~/Library/LaunchAgents/com.zettelmerken.dailyreview.plist`

1. Unload agent `launchctl disable ~/Library/LaunchAgents/com.zettelmerken.dailyreview.plist`

1. Start/Stop (for debugging) `launchctl start com.zettelmerken.dailyreview`

1. List all services `launchctl list`

## Error: Unrecognized target specifier.

- Add a target specifier

```bash
launchctl enable user/$(id -u)/com.abc.xyz.plist
```
