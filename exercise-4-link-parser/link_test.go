package main

import (
	"reflect"
	"testing"
)

func TestGetLinksFromHtml(t *testing.T) {
	tests := []struct {
		name      string
		filename      string
		wantLinks []Link
	}{
		{"Get single link with text", "example1.html", []Link{{"/other-page", "A link to another page"}} },
		{"Get multiple links, excluding html within <a>", "example2.html", []Link{
			{"https://www.twitter.com/joncalhoun", "Check me out on twitter"},
			{"https://github.com/gophercises", "Gophercises is on Github!"},
		}},
		{"Get multiple links from a larger file, excluding html within <a>", "example3.html", []Link{
			{"#", "Login"},
			{"/lost", "Lost? Need help?"},
			{"https://twitter.com/marcusolsson", "@marcusolsson"},
		}},
		{"Get single link excluding child comment", "example4.html", []Link{{"/dog-cat", "dog cat"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotLinks := GetLinksFromHtml(parseHtmlFile(tt.filename)); !reflect.DeepEqual(gotLinks, tt.wantLinks) {
				t.Errorf("GetLinksFromHtml() = %v, want %v", gotLinks, tt.wantLinks)
			}
		})
	}
}
