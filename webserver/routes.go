package webserver

import (
	"fmt"
	"io"
	"net/http"

	"github.com/John-Holden/goBlog/core"
	"github.com/gorilla/mux"
)

func (b *Blog) root(w http.ResponseWriter, r *http.Request) {
	head := core.GetHead(
		b.headTitle,
		b.HeadDescription,
		b.static+"/favicon.ico",
		[]string{b.static + "/default.css"},
		[]string{})

	body := core.GetPostListBodyHtml(b.content)
	page := core.GetPage(head, body)
	io.WriteString(w, page)
}

// Renders a single post
func (b *Blog) post(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[i] Rendering Post: " + r.URL.Path)
	html_doc := ""
	links := core.GetPostPaths(b.content)
	for link, filename := range links {
		if "/"+b.content+"/"+link == r.URL.Path {
			html_doc = core.GetPostHtml(b.content + "/" + filename)
			break
		}
	}

	if html_doc == "" {
		http.Error(w, "404 not found.", http.StatusNotFound)
	}

	io.WriteString(w, html_doc)
}

func SetRoutes(blog Blog) *mux.Router {
	fmt.Println("[i] Setting WebServer routes...")

	r := mux.NewRouter()
	r.HandleFunc("/", CorsMiddleware(blog.root))
	r.HandleFunc("/"+blog.content+"/{post}", CorsMiddleware(blog.post))
	r.HandleFunc("/"+blog.static+"/{css}", CSSMiddleware(ServeCSS))
	r.PathPrefix(blog.static).Handler(
		http.StripPrefix(
			blog.static,
			http.FileServer(
				http.Dir(blog.static),
			),
		),
	)
	return r
}
