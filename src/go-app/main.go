package main

import(
	"fmt"
	"net/http"

	"google.golang.org/appengine" // Required external App Engine Library
)
func indexHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	fmt.Fprintln(w, "Hello, World! BIATCH")
}

func main(){
	http.HandleFunc("/", indexHandler)
	appengine.Main()
}
