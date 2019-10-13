# texfmt [![Build Status](https://travis-ci.org/andykuszyk/texfmt.svg?branch=master)](https://travis-ci.org/andykuszyk/texfmt)
`texfmt` is a simple line width formatter for LaTeX files. Its purpose is to reformat `tex` files to a fixed width in a similar way to existing programs like `fold`.

Unlike `fold`, however, `texfmt` is aware of LaTeX syntax (to some degree!) and is able to concatenate lines (as well as splitting them like `fold`) accordingly.

## Usage
```
texfmt -w 120 path/to/my_file.tex
```

Will reformat `my_file.tex` to a maximum width of 120 characters and output the new file to the standard output.

## Installation
* If you have `go` installed: `go get github.com/andykuszyk/texfmt`
* Download the latest release from the releases page.

## Releasing
Tag a commit on `master` and `git push --tags`.
