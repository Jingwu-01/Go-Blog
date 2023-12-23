package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleFunc)
	http.ListenAndServe(":3000", nil)
}

func handleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	fmt.Println(r.URL.Path)
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1> Hello, this is go blog. </h1>")
	} else if r.URL.Path == "/about" {
		fmt.Fprint(w, "<p>This is a blog that records my programming notes. Please contact me at <a href=\"mailto:123@rice.edu\">123@rice.edu</a> </p>")
	} else {
		fmt.Fprint(w, "<h1> We did not find your request.</h1>")
	}
}
