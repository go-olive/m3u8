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
