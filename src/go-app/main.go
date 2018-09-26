package main

import(
	"net/http"

	"google.golang.org/appengine" // Required external App Engine Library
)
func indexHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
}

func main(){
	http.HandleFunc("/", indexHandler)
	appengine.Main()
}
