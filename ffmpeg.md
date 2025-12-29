---
modified: "Fri Dec 19 14:06:37 EST 2025"
---

# ffmpeg

## Convert a mov to mp4

```sh
ffmpeg -i input.mov output.mp4
```

## Trim media from/to

```sh
ffmpeg -ss <start_seconds> -i "input.mp3" -t <end_second> "output.mp3"
```

## Compress a video

```bash
ffmpeg -i input.mp4 -c:v libx265 -crf 25 output.mp4
```
