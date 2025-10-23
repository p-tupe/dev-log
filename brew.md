---
modified: Wed Nov 23 17:59:51 EST 2022
---
# brew

> The Missing Package Manager for macOS

- [Official Site](https://brew.sh/)

### Essential Commands

1. General help: `brew --help`

1. Search for a package: `brew search <package>`

1. Get info about a package: `brew info <package>`

1. Install a package (formula or cask): `brew install [--no-quarantine] [--cask] <package>`

   - I normally don't use the optional `--cask` flag, brew is clever enough to know which I want most of the time; and prompts me for which when ambiguous.

   - `--no-quarantine` ensures you don't get the security confirmation dialog on opening a cask for first time.

   - `formula` refers to a cli app or lib while `cask` refers to a full "gui" application

1. Remove a package: `brew uninstall <package>`

1. Update/Upgrade all packages: `brew update && brew upgrade [--greedy]`

   - `--greedy` flag is for those apps that have auto-update enabled. I use it always, instead of upgrading individual app.

1. Cleanup everything: `brew autoremove && brew cleanup [--prune=all]`

1. List of all formulae explicitly installed (without deps): `brew leaves`
