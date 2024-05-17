package webserver

type Blog struct {
	headTitle,
	HeadDescription,
	content,
	static,
	port string
}

func SetBlog(
	content_dir,
	static_dir,
	port string) *Blog {
	return &Blog{
		headTitle:       "Blog List",
		HeadDescription: "My blog list",
		content:         content_dir,
		static:          static_dir,
		port:            port,
	}
}
