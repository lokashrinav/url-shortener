package main

import (
	"fmt"
	"net/http"
)


func main() {
	mux := http.NewServeMux()

	smth := map[string]string {
		"/url": "https://gophercises.com/",
		"/dog": "https://github.com/lokashrinav?tab=repositories",
	}

	mux.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		
	})

	http.ListenAndServe(":3000", mux)
	
}
