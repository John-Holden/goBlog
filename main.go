package main

import (
	"github.com/John-Holden/goBlog/webserver"
)

func main() {
	content_dir := "content"
	static_dir := "static"
	port := "8000"

	blog := webserver.SetBlog(content_dir, static_dir, port)
	webserver.ServeLocal(*blog)
}
