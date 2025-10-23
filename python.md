---
modified: Mon Jul 18 18:32:52 EDT 2022
---
# Python

## Links

- [Official site](https://www.python.org/)

- [Official docs](https://docs.python.org/3/tutorial/index.html)

- [Standard Lib](https://docs.python.org/3/py-modindex.html)

- [Hitchhiker's Guide to Python: Online Docs](https://docs.python-guide.org/)

- [Exercism site: Guided Exercises](https://exercism.org/tracks/python)

- [Awsome python collection github](https://github.com/vinta/awesome-python)

- [Automate The Boring Stuff: Book](https://automatetheboringstuff.com/2e/)

- [Think Python: Book](https://www.greenteapress.com/thinkpython/html/index.html)

- [Practical Python: Series of Article](https://dabeaz-course.github.io/practical-python/)

- [Snarky Canadian: Blog series of python](https://snarky.ca/)

- [Youtube: What does it take to be an expert in python](https://www.youtube.com/watch?v=7lmCu8wz8ro)

- [Pep8: Code formatting guideline](https://www.python.org/dev/peps/pep-0008)

- [Python Language Reference](https://docs.python.org/3/reference/index.html)

### Installing packages in a Virtual Environment

- [Source](https://snarky.ca/a-quick-and-dirty-guide-on-how-to-install-packages-for-python/)

  ```python
  python -m venv --prompt . .venv
  echo "*" > .venv/.gitignore
  source .venv/bin/activate
  python -m pip install --upgrade pip
  ```

  To deactivate: `source .venv/bin/deactivate`
