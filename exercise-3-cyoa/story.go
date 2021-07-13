package main

import (
	"encoding/json"
	"html/template"
	"log"
)

type story struct {
	Title string `json:"title"`
	Story []string `json:"story"`
	Options []option `json:"options"`
}

type option struct {
	Text string `json:"text"`
	Arc string `json:"arc"`
}

func parseStories(storyBytes []byte) (stories map[string]story){
	err := json.Unmarshal(storyBytes, &stories)

	if err != nil {
		log.Fatal(err)
	}

	return stories
}

func buildTemplateForStory() *template.Template{
	htmlTemplate, err := template.New("StoryTemplate").Parse(string(parseFile("template.html")))

	if err != nil {
		log.Fatal(err)
	}

	return htmlTemplate
}
