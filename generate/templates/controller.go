package templates

// Controller
var Controller = `package controllers

import "github.com/gophersaurus/gf.v1/http"

// {{ . }} is a controller.
var {{ . }} = struct {
	Index func(resp http.Responder, req *http.Request)
}{
	Index: func(resp http.Responder, req *http.Request) {

		// write the result
		resp.WriteFormat(req, req)
	},
}`
