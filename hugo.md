---
modified: Mon Apr 18 21:58:25 EDT 2022
---
# HUGO

## Links

- [Official Site](https://www.hugotutorial.com/posts/)

- [Blogs](https://www.hugotutorial.com/posts/)

## Commands

- Install: `yay -S hugo`

- Help: `hugo help`

- New Site: `hugo new site <site-name>`

- New Post: `hugo new posts/<post-name>.md`

  - Dir `post` is not mandatory, can be changed

  - Adds file with date-name.md and some front-matter

- Build: `hugo`

- Serve: `hugo server [options]`
  - -D for drafts enabled
  - -p <port>
  - --navigateToChanged

## Directory structure

```
/hugo-site
├── archetypes
├── content
├── data
├── layouts
└── static
```

- archetypes: `hugo new <archetype>` Set custom front-matter and content (markdown)

- content: meat of site (markdown)

  - Also determines relative urls and layout used depending on dir name and structure

- data: lotsa shared front-matter of sorts (json, yaml, toml)

- layouts: templates for content (html)

- static: (img, css, js blah)

## hugo without theme

Hugo pulls data from `content/`, `data/`, `static/` and structures it by templates defined in `layouts/` dir.

Content -> Layout mapping: [Template Lookup Order](https://gohugo.io/templates/lookup-order/)

NOTE: Leaving in favor of [11ty](./11ty.md), couldn't get a single page up without writing 4 more files and skimming docs and posts for hours. Not what I am looking for.
