---
modified: Wed Jun  1 21:19:07 EDT 2022
---
# commonlisp

- [Official Site](https://common-lisp.net/)

https://lispcookbook.github.io/cl-cookbook/getting-started.html

https://cliki.net/

https://gigamonkeys.com/book/

https://stevelosh.com/blog/2018/08/a-road-to-common-lisp/

https://learnxinyminutes.com/docs/common-lisp/

https://exercism.org/tracks/common-lisp/

http://www.lispworks.com/documentation/HyperSpec/Front/Contents.htm

IMPORTANT: I will be using SBCL implementation, so any notes that follow will be for same

## Notes

Installing quicklist (a package manager)

```shell
$ curl -O https://beta.quicklisp.org/quicklisp.lisp
$ sbcl --load quicklisp.lisp
* (quicklisp-quickstart:install)
* (ql:system-apropos "vecto")
* (ql:quickload "vecto")
* (ql:add-to-init-file)
* (quit)
```

Making REPL usable (since no history or completion by default)

```shell
$ rlwrap -c sbcl --disable-debugger
```

> If using neovim (as you should be), use Iron.nvim.
> It provides a nice repl integration for _many_ languages, including lisp.
