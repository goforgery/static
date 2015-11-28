# static

[![Build Status](https://secure.travis-ci.org/goforgery/static.png?branch=master)](http://travis-ci.org/goforgery/static)

Static file server for Forgery2.

## Install

	go get github.com/goforgery/static

## Use

By default files are served from a folder named `./public` in the directory where __Forgery2__ is started. For example if you have a the file `./public/text.txt` and you start the server in `./`, the file will be served from the URL http://localhost:3000/text.txt.

```javascript
package main

import (
	"github.com/goforgery/forgery2"
	"github.com/goforgery/static"
)

func main() {
	app := f.CreateApp()
	app.Use(static.Create())
	app.Listen(3000)
}
```

## Test

    go test
