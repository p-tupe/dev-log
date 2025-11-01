---
modified: "Fri Oct 31 20:18:59 EDT 2025"
---

## ffmpeg

- Convert a mov to mp4 (with compression):

```sh
ffmpeg -i video.mov -vcodec h264 -acodec mp2 video.mp4
```

- Trim media from/to

```sh
ffmpeg -ss <start_seconds> -i "input.mp3" -t <end_second> "output.mp3"
```

- Compress a video

```bash
ffmpeg -i input_video.mp4 -c:v libx265 -crf 25 output_video.mp4
```
