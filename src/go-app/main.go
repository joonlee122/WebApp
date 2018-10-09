package main

import(
	"fmt"
	"google.golang.org/appengine" // Required external App Engine Library
	"github.com/go-fsnotifiy/fsnotify"
	
)

func main(){
	// creates a new file watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("ERROR", err)
	}
	defer watcher.Close()

	//
	done := make(chan bool)

	//
	go func() {
		for {
			select {
			// watch for events
			case event := <-watcher.Events:
				fmt.Printf("EVENT! %#v\n", event)

				// watch for errors
			case err := <-watcher.Errors:
				fmt.Println("ERROR", err)
			}
		}
	}()

	// out of the box fsnotify can watch a single file, or a single directory
	if err := watcher.Add("/Users/ttoya/Desktop/XponentialWorks"); err != nil {
		fmt.Println("ERROR", err)
	}

	<-done
}
