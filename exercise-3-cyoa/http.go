package main

import (
	"fmt"
	"net/http"
)

func startServer() {
	fmt.Println("Listening on port 8080")

	storyHandler := serveStories([]story{}, defaultMux())
	http.ListenAndServe(":8080", storyHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	return mux
}

func indexHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "Choose your own adventure")
}

func serveStories(stories []story, fallback http.Handler) http.HandlerFunc {
	return func (response http.ResponseWriter, request *http.Request) {
		fallback.ServeHTTP(response, request)
	}
}
