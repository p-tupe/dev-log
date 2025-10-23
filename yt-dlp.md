---
modified: Thu Oct 23 12:48:34 EDT 2025
---
## Youtube Download

- Only audio:

  ```sh
  yt-dlp --ignore-errors --format bestaudio --extract-audio --audio-format mp3 --output '%(title)s.%(ext)s' [--yes-playlist] <"URL" | --batch-file  ./fileName>
  ```

  _where fileName is a list of urls separated by blank lines;_

- Best video+audio:

  ```sh
  yt-dlp --ignore-errors --format 'bestvideo[ext=mp4]+bestaudio[ext=m4a]/bestvideo+bestaudio' --merge-output-format mp4  -o '%(title)s.%(ext)s' <"video-link">
  ```

  Note the brackets for the above are part of the command.
