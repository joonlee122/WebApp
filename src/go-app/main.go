package main

import(
	"net/http"
	"fmt"
	"html/template"
	"google.golang.org/appengine" // Required external App Engine Library
)
func indexHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	params := templateParams{}

	if r.Method == "GET" {
		indexTemplate.Execute(w, params)
		return
	}

	//It's a POST request, so handle the form submission

	name := r.FormValue("name")
	params.Name = name
	if name == "" {
		name = "Anonymous Gopher"
	}

	if r.FormValue("message") == "" {
		w.WriteHeader(http.StatusBadRequest)

		params.Notice = "No message provided"
		indexTemplate.Execute(w, params)
		return
	}

	params.Notice = fmt.Sprintf("Thank you for your submission, %s!", name)
}

var (
	indexTemplate = template.Must(template.ParseFiles("index.html"))
)

type templateParams struct {
	Notice string
	Name string
}

func main(){
	http.HandleFunc("/", indexHandler)
	appengine.Main()
}
