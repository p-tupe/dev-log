# (NEO)VIM

- Redirect output of Ex commands to register

1. `:redir @a`
2. `:__commands__`
3. `:redir END`
4. `"ap`

- Redirect output of telesope to quickfix `C-q`

- Run bash (or any shell) for buffer content

```
:%!bash
```

- Convert textfile to pdf (via ps)

```sh
$ vim filename.txt -c "hardcopy > filename.ps | q"; pstopdf filename.ps
```

_pstopdf may be names ps2pdf on other OSes_

- Alternative undo redo

```vim
:earlier __time__
:later __time__
```

where **time** = x(s|m|h|d) (s = second, m = minute, h = hour, day)

to go back and forth on state of buffer in time of x units

- Remote editing (using ssh config)

```vim
:e scp://staging/~/path/to/file.ext
```

_Assuming a Host for 'staging' is set in ~/.ssh/config_

- Delete other buffers except current

```vim
:% bd | e # | bd #
```

where

- % bd = delete all buffers (creates a new no-name buffer),
- e # = edit previous buffer (since current is no-name),
- bd # = delete previous buffer (no-name)

_spaces are optional_

- Ex mode completion options

TAB cycles through options; Ctrl-a inputs all options.

Example: To delete all \*.js files in buffer list

```vim
:bd js Ctrl-a
```

_Can use :bwipeout to completely remove a buffer_

- Search and Replace (Using neovim telesope)

  - Search for a term `:Telescope live_grep`
  - Put all results in quickfix buffer `<C-q>`
  - Run replace interactive on all results `:cdo s/<sterm>/<rterm>/g[ci] | update`
  - Capture groups are defined by `\1` `\2` etc and original string by `&`

- Check diff of files

  ```shell
  $ nvim -d -R file_one file_two
  ```

  - -d = diff
  - -R = Readonly
