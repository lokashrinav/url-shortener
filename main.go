package main

import (
	"net/http"
)


func main() {
	mux := http.NewServeMux()

	urlMap := map[string]string {
		"/url": "https://gophercises.com/",
		"/dog": "https://github.com/lokashrinav?tab=repositories",
		"/go":       "https://golang.org",
		"/openai":   "https://openai.com",
		"/github":   "https://github.com",
	}

	mux.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		value, ok := urlMap[path]; 
		if ok {
			http.Redirect(w, r, value, http.StatusFound)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})

	http.ListenAndServe(":8080", mux)
	
}
