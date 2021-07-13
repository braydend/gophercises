package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	flags := setupFlags()
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yamlHandler, err := YAMLHandler(readFile(flags.yamlFile), mapHandler)
	if err != nil {
		panic(err)
	}
	jsonHandler, err := JSONHandler(readFile(flags.jsonFile), yamlHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", jsonHandler)
}

type Flags struct {
	yamlFile string
	jsonFile string
}

func setupFlags() Flags {
	var yamlFile string
	var jsonFile string
	flag.StringVar(&yamlFile, "yaml", "default.yml", "Yaml file to read path/url pair from")
	flag.StringVar(&jsonFile, "json", "", "Yaml file to read path/url pair from")

	flag.Parse()

	return Flags{yamlFile, jsonFile}
}

func readFile(filename string) []byte{
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	bytes, err := ioutil.ReadAll(file)

	if err != nil {
		log.Fatal(err)
	}

	return bytes
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
