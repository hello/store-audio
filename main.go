package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Handler struct {
	audioPath string
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Println("HTTP POST is required")
	}

	defer r.Body.Close()
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	now := time.Now().Format("2006-01-02-150405")
	ioutil.WriteFile(now, content, 0777)
	// fmt.Fprintf(w, "File saved: %s\n", now)

	audio, err := ioutil.ReadFile(h.audioPath)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "%s", audio)
}

func main() {

	port := flag.String("port", "4567", "port to bind local server to")
	audioPath := flag.String("audio", "audio.raw", "path to audio file")

	flag.Parse()
	h := Handler{
		audioPath: *audioPath,
	}
	http.Handle("/", h)
	http.ListenAndServe(":"+*port, nil)
}
