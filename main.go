package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", defaultHandler)
	router.HandleFunc("/about", aboutHandler)
	router.HandleFunc("/articles/", articlesHandler)

	http.ListenAndServe(":3000", router)
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1> Hello, this is the updated go blog. </h1>")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1> We did not find your request.</h1>")
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/about" {
		fmt.Fprint(w, "<p>This is a blog that records my programming notes. Please contact me at <a href=\"mailto:123@rice.edu\">123@rice.edu</a> </p>")
	}
}

func articlesHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.SplitN(r.URL.Path, "/", 3)[2]
	fmt.Fprint(w, "Article ID:"+id)
}
