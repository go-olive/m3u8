# OliveM3u8

[![GoDoc](https://img.shields.io/badge/GoDoc-Reference-blue?style=for-the-badge&logo=go)](https://pkg.go.dev/github.com/go-olive/m3u8?tab=doc)

OliveM3u8 is a golang library which downloads m3u8 video streams.

## API Guide

```go
package main

import (
	"github.com/go-olive/m3u8"
)

func main() {
	p := m3u8.NewParser(&m3u8.Config{
		// m3u8 url
		Url: "https://video/index.m3u8",
		// video name
		Out: "video",
		// video directory
		Dir: "/",
	})
	p.Parse()
}

```

## Credits

This project is inspired by [m3u8-downloader](https://github.com/llychao/m3u8-downloader)