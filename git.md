# Git

## Stash

> Stash the changes in a dirty working directory away

- Show all saved points: `git stash list`

- Quickly stash current changes: `git stash [push]`

- Quickly unstash latest save: `git stash pop`

- Clear all saved stashes: `git stash clear`

- Save with a specific name: `git stash save <name>`

- Apply a specific stash: `git stash apply <n>`

  - where "n" is the stash@{n} you see after listing all stashes

## Rebase

> Make it look like your commit(s) were made on top of the latest head

- In the branch to be rebased: `git rebase <rebase-on>`

- Example, to rebase a featurex branch on main: `git checkout featurex; git rebase main; git push -f origin featurex`

- <span style="color: red">IMPORTANT!</span> Do not rebase on shared public branches!

### Manual Rebase

- If changed committed in latest commit: `git reset --soft HEAD~1`

- Stash working state: `git stash`

- Pull latest from upstream: `git pull origin ...`

- Pop working state: `git stash pop`

- Create a commit: `git commit`

- Force push: `git push -f origin ...`

## Tags

> Mark a point in commit history (useful for releases and such)

- [Docs](https://git-scm.com/book/en/v2/Git-Basics-Tagging)

- Create a tag (annotated)

  `git tag -a v0.2 -m "Release v0.2"`

  Note that simple `git push` does not push a tag to remote, must do something like:

  `git push origin <tag>`

- List tags

  `git tag -l`

- Show specific tag details

  `git show v0.2`

- Delete all branches except main and featurex

  ```sh
  git branch | awk '!/main|featurex/ { print $1 }' | xargs git branch -D
  ```

## Submodules

> Make a repository a subdirectory of another repository

- [Docs](https://git-scm.com/book/en/v2/Git-Tools-Submodules)

- Add a submodule

  ```sh
  git submodule add <sub-repo-link>
  ```

  Creates a .gitmodules file with path and url of submodules and adds a submodule in detached HEAD state

- Clone with submodules

  ```sh
  git clone --recurse-submodules <main-repo-link>
  ```

- If already clone but no submodules, do

  ```sh
  git submodule update --init --recursive
  ```

- Update submodules from remote

  ```sh
  git submodule update --remote [name]
  ```

- The 'foreach' command

  ```sh
  git submodule foreach 'git pull'
  ```

## Worktrees

> Make a directory for each working branch

- Useful when working on multiple branches, to avoid stash/pop unmerged paths confusion and work without disturbing other branches. Parallel branches ftw!

- Having multiple folders for each branch makes things easy for IDEs, compared to restructuring same folder multiple times

- <span style="color:orange">IMPORTANT!</span> Do not use with submodules

- Not essential but a better workflow is to start off with a bare repository and have one folder each for each branch:

  ```sh
  mkdir <repo>; cd <repo>
  git clone --bare <remote>
  echo "gitdir: ./." > .git
  git config remote.origin.fetch "+refs/heads/*:refs/remotes/origin/*"
  git fetch origin
  git worktree add <branch>
  ```

- Worktree commands help: `git worktree help`
