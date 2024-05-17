package webserver

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Serve static CSS files
func ServeCSS(w http.ResponseWriter, r *http.Request) {
	cssFilePath := r.URL.Path[1:]
	cssContent, err := ioutil.ReadFile(cssFilePath)
	if err != nil {
		fmt.Println("Error reading CSS file:", err)
		return
	}
	io.WriteString(w, string(cssContent))
}

func ServeLocal(blog Blog) {
	r := SetRoutes(blog)
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:" + blog.port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
