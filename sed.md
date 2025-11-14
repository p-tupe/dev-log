---
modified: "Thu Nov 13 16:30:49 EST 2025"
---

# sed (stream editor)

## Add an extra line for each line break in a file:

```sh
echo "Line 1\nLine 2" | sed "G"
# --> "\nLine 1\n\nLine 2\n"
```

## Prepend some output to the start of a file:

```bash
# Say we wanna add a file.md's last modified time to it's frontmatter
modified=$(date -r file.md); sed -e "1s/^/---\nmodified: "$modified"\n---\n/g" file.md
```

## Update a specific line number

Notice the `1` in previous command? It represents a line number. So, let update the "modified" time of the file on 2nd line:

```bash
sed "2 s/.*/modified: \"$(date -r file.md)\"/" file.md
```

where `2` in the command represents the second line

## Or a range of lines

```bash
sed -n "1,5p" file.md
```

To show the first 5 lines of `file.md`
