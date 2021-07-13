package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func startServer(stories map[string]story) {
	storyHandler := serveStories(stories, defaultMux())

	fmt.Println("Listening on port 8080")
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

func getPathWithoutSlash(url *url.URL) string {
	return strings.Replace(url.Path, "/", "", 1)
}

func serveStories(stories map[string]story, fallback http.Handler) http.HandlerFunc {
	return func (response http.ResponseWriter, request *http.Request) {

		tmpl := buildTemplateForStory()

		story, exists := stories[getPathWithoutSlash(request.URL)]

		if !exists {
			story = stories["intro"]
		}

		err := tmpl.ExecuteTemplate(response, "StoryTemplate", story)

		if err != nil {
			log.Fatal(err)
		}
	}
}
