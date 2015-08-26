# Atexit

Package mailtoweb returns the webmail URL related to a an email address.
It can be used in web apps to provide a direct link to the user's webmail from his email address, after a subscription for example.

## Installation

```Shell
$ go get github.com/whitedevops/mailtoweb
```

## Usage [![GoDoc](https://godoc.org/github.com/whitedevops/mailtoweb?status.svg)](https://godoc.org/github.com/whitedevops/mailtoweb)

```Go
package main

import (
	"github.com/whitedevops/mailtoweb"
)

func main() {
	mailtoweb.Use("user@gmail.com")
}
```