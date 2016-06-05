package main

import (
	"countWords/wordCore"
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	Text string
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome, % !", r.URL.Path[1:])
}

func wordCount(w http.ResponseWriter, r *http.Request) {
	b1, err2 := json.Marshal(wordCore.SingleWordCount(r.URL.Query().Get("sentence"), true))
	if err2 != nil {
		panic(err2)
	}
	w.Write(b1)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/wordCount/", wordCount)
	http.ListenAndServe(":8080", nil)
}
