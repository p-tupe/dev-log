---
modified: "Fri Nov  7 22:07:37 EST 2025"
---

# Youtube Download

## Download best audio

```sh
yt-dlp --ignore-errors --format bestaudio \
--extract-audio --audio-format mp3 \
--output '%(title)s.%(ext)s' [--yes-playlist] <video-link | --batch-file  ./file.txt>
```

where file.txt contains a list of urls separated by blank lines.

## Download best audio+audio:

```sh
yt-dlp --verbose --ignore-errors \
  --format 'bestvideo+bestaudio' \
  --merge-output-format mp4  \
  --output '%(title)s.%(ext)s' <video-link>
```
