# csplit
modified: "Sun Oct 26 21:38:59 EDT 2025"
> split a big file (or stream) into smaller chunks based on pattern

Say we have a file like so:

```txt
This is the first section.
---
This is another section.
---
This is a final section.
```

To split it (on MacOS using default `csplit`), we do:

```bash
csplit file.txt /---/ {1}
```

which will result in 3 files `xx00`, `xx01` and `xx02` with contents:

```bash
# xx00
This is the first section.

# xx01
---
This is another section.

# xx02
---
This is a final section.

```

Here `{1}` is the increment count of splits to do. If we omit it, it'll split the file in twain.

> See `man csplit` for more options

On GNU coreutils, we may supply count as "{\*}" for however many splits required, but here we must pass a number (which can be arbitrarily large).
