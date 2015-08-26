package stackr

import (
	"github.com/goforgery/forgery2"
	. "github.com/ricallinson/simplebdd"
	"testing"
)

func TestCreate(t *testing.T) {

	Describe("Create()", func() {

		var app *f.Application
		var req *f.Request
		var res *f.Response

		BeforeEach(func() {
			app = f.CreateApp()
			req = f.CreateRequestMock(app)
			res, _ = f.CreateResponseMock(app, false)
		})

		It("should return [404]", func() {
			app.Use("", Create())
			app.Handle(req, res, 0)
			AssertEqual(res.StatusCode, 404)
		})

		It("should return [404]", func() {
			app.Use("", Create(map[string]string{}))
			app.Handle(req, res, 0)
			AssertEqual(res.StatusCode, 404)
		})

		It("should return [404] from a directory", func() {
			req.OriginalUrl = "/directory/"
			app.Use("", Create(map[string]string{"root": "./fixtures/"}))
			app.Handle(req, res, 0)
			AssertEqual(res.StatusCode, 404)
		})

		It("should return [404] from a directory without a trailing slash", func() {
			req.OriginalUrl = "/directory"
			app.Use("", Create(map[string]string{"root": "./fixtures/"}))
			app.Handle(req, res, 0)
			AssertEqual(res.StatusCode, 404)
		})

		It("should return [200] from a file", func() {
			req.OriginalUrl = "/text.txt"
			app.Use("", Create(map[string]string{"root": "./fixtures/"}))
			app.Handle(req, res, 0)
			AssertEqual(res.StatusCode, 200)
		})

		It("should return [404] from a directory on /public path", func() {
			req.OriginalUrl = "/public/directory/"
			app.Use("/public", Create(map[string]string{"root": "./fixtures/"}))
			app.Handle(req, res, 0)
			AssertEqual(res.StatusCode, 404)
		})

		It("should return [200] from a file on /public path", func() {
			req.OriginalUrl = "/public/text.txt"
			app.Use("/public", Create(map[string]string{"root": "./fixtures/"}))
			app.Handle(req, res, 0)
			AssertEqual(res.StatusCode, 200)
		})

		It("should return [404] from a directory", func() {
			req.OriginalUrl = "/directory/"
			app.Use("", Create(map[string]string{"root": "./fixtures"}))
			app.Handle(req, res, 0)
			AssertEqual(res.StatusCode, 404)
		})

		It("should return [200] from a opt.Root with no trailing slash", func() {
			req.OriginalUrl = "/text.txt"
			app.Use("", Create(map[string]string{"root": "./fixtures"}))
			app.Handle(req, res, 0)
			AssertEqual(res.StatusCode, 200)
		})

		It("should return [200] from a cached file", func() {
			req.OriginalUrl = "/text.txt"
			app.Use("", Create(map[string]string{"root": "./fixtures"}))
			app.Handle(req, res, 0)
			app.Handle(req, res, 0)
			app.Handle(req, res, 0)
			AssertEqual(res.StatusCode, 200)
		})
	})

	Report(t)
}
