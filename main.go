package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"text/template"

	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, this is the updated go blog. </h1>")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<p>This is a blog that records my programming notes. Please contact me at <a href=\"mailto:123@rice.edu\">123@rice.edu</a> </p>")

}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "<h1> Page is not found </h1>")
}

func articlesShowHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintf(w, "Article ID:"+id)
}

func articlesIndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get Article Successful")
}

func articlesStoreHandler(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	body := r.FormValue("body")
	errors := make(map[string]string)

	if title == "" {
		errors["title"] = "Title cannot be empty"
	} else if len(title) < 3 || len(title) > 40 {
		errors["title"] = fmt.Sprintf("The length of title should be between 3 to 40 characters. The current length is %v", len(title))
	}

	if body == "" {
		errors["body"] = "Content cannot be empty"
	} else if len(body) < 10 {
		errors["body"] = fmt.Sprintf("The length of body should be greater or equal to 10 characters. The current length is %v", len(body))
	}

	if len(errors) == 0 {
		fmt.Fprintf(w, "Validation passed! <br>")
		fmt.Fprintf(w, "The value of title is %v <br>", title)
		fmt.Fprintf(w, "THe value of body is %v <br>", body)
	} else {
		storeURL, _ := router.Get("articles.store").URL()

		data := ArticlesFormData{
			Title:  title,
			Body:   body,
			URL:    storeURL,
			Errors: errors,
		}
		tmpl, err := template.ParseFiles("resources/views/articles/create.gohtml")
		if err != nil {
			panic(err)
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			panic(err)
		}
	}

}

func articlesCreateHandler(w http.ResponseWriter, r *http.Request) {
	storeURL, _ := router.Get("articles.store").URL()
	data := ArticlesFormData{
		Title:  "",
		Body:   "",
		URL:    storeURL,
		Errors: nil,
	}

	tmpl, err := template.ParseFiles("resources/views/articles/create.gohtml")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		panic(err)
	}
}

type ArticlesFormData struct {
	Title, Body string
	URL         *url.URL
	Errors      map[string]string
}

func HTMLMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		h.ServeHTTP(w, r)
	})
}

func removeTrailingSlash(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		}
		h.ServeHTTP(w, r)
	})
}

func main() {

	router.HandleFunc("/", homeHandler).Methods("GET").Name("home")
	router.HandleFunc("/about", aboutHandler).Methods(("GET")).Name("about")

	router.HandleFunc("/articles/{id:[0-9]+}", articlesShowHandler).Methods("GET").Name("articles.show")
	router.HandleFunc("/articles", articlesIndexHandler).Methods("GET").Name("articles.index")
	router.HandleFunc("/articles", articlesStoreHandler).Methods("POST").Name("articles.store")
	router.HandleFunc("/articles/create", articlesCreateHandler).Methods("GET").Name("articles.create")

	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	router.Use(HTMLMiddleware)

	http.ListenAndServe(":3000", removeTrailingSlash(router))
}
