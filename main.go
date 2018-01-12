package main

import (
	"io/ioutil"
	"net/http"
	"os"
)

func binaryHandler(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("output/fake")
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-type", "application/octet-stream")
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	w.Write(bytes)
}

func textHandler(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("output/shell.sh")
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-type", "text/plain")
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	w.Write(bytes)
}

func main() {
	http.HandleFunc("/commands/foo/binary/1.0.2", binaryHandler)
	http.HandleFunc("/commands/foo/text/1.0.1", textHandler)
	http.ListenAndServe(":8080", nil)
}
