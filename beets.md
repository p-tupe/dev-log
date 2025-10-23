## beets

- Install

```bash
uv tool install  beets  --with requests --with chromaprint --with pyacoustid
```

- Seek help

  ```shell
  $ beet help [command]
  ```

- For all tracks in albumFolder, import metadata, fix filenames and copy to Dir:

  ```shell
  $ beet -d ~/Dir import [-S <url>] [-s <file> | -g <folder>]
  ```

  where  
  \-S - Source URL (musicbrainz)
  \-s - Import as single track (Artist -> Track)
  \-g - Import as album (Album -> Track)
