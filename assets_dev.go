// +build dev
//go:generate go run assets_generate.go

package gohttpserver

import "net/http"

// Assets contains project assets.
var Assets http.FileSystem = http.Dir("assets")
