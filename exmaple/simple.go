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
