package main

import (
	"encoding/json"
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if pathsToUrls[request.RequestURI] != "" {
			http.Redirect(writer, request, pathsToUrls[request.RequestURI], http.StatusPermanentRedirect)
		} else {
			fallback.ServeHTTP(writer, request)
		}
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
type pathUrl struct {
	Path string `yaml:"path" json:"path"`
	Url string `yaml:"url" json:"url"`
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var pathUrls []pathUrl
	err := yaml.Unmarshal(yml, &pathUrls)

	if err != nil {
		return nil, err
	}

	pathToUrls := make(map[string]string)

	for _, data := range pathUrls {
		pathToUrls[data.Path] = data.Url
	}

	return MapHandler(pathToUrls, fallback), nil
}

func JSONHandler(jsonBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var pathUrls []pathUrl
	err := json.Unmarshal(jsonBytes, &pathUrls)

	if err != nil {
		return nil, err
	}

	pathToUrls := make(map[string]string)

	for _, data := range pathUrls {
		pathToUrls[data.Path] = data.Url
	}

	return MapHandler(pathToUrls, fallback), nil
}