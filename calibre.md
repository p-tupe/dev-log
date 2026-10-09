---
modified: 2026-10-09 16:05:01.047274 -0400 EDT
---
# calibre

## How I use it

I don't use the app in the way it's supposed to be used, I think. For me, it's more of a tool to enhance the ebooks (add cover & metadata, renam file etc) than keep an organized library.

Here's the commands I've setup so my ebook collection is up to sniff:

1. Dump an ebook (generally `.pdf`) into my `~/Ebooks` folder

2. Check for info already embedded inside:

```bash
ebook-meta book.pdf
```

> It's best to grab `ISBN`, should be in there somewhere

3. Run the following shell commands to shelf it well:

```bash
fetch-ebook-metadata --isbn ".." --title ".." --authors ".." -o book.opf && \
ebook-meta book.pdf --from-opf book.opf
```

> --authors separates multiple authors with &, not commas

And that's all! Synced everywhere as a proper book.

---

Here's a full script that processes each file in the folder automatically:

```bash
# todo
```

dump this as `update-ebooks` in your `bin` and write a `launchd`/`systemd` script to run it on `~/Ebooks` update and Voîla! Automatic ebook management!
