---
modified: Thu Oct 23 18:34:54 EDT 2025
---
## sed (stream editor)

- Add an extra line for each line break in a file:

```sh
echo "Line 1\nLine 2" | sed "G"
# --> "\nLine 1\n\nLine 2\n"
```

- Prepend some output to the start of a file:

```bash
# Say we wanna add a file.md's last modified time to it's frontmatter
modified=$(date -r file.md); sed -e "1s/^/---\nmodified: "$modified"\n---\n/g" file.md
```
