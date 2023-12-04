package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	// responseWriter provides methods to assemble http response and send to user
	// *http.Request is a pointer to a struct that holds info about current request,
	// like http method and url being requested

	// serve 404 if not exact path, this way we avoid using it as a catchall
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display specific snippet..."))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		// w.WriteHeader(405)
		// w.Write([]byte("Method not allowed"))
		// or can use the following
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Creating a new snippet"))
}

func main() {
	fmt.Println("hi!")

	mux := http.NewServeMux()
	// in servemux / is a catch all, all http requests will be routed to this
	// this is because a subtree path (path with trailing /) will match anything
	// after it e.g. like a wildcard /* /snippet/= /snippet/*
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("sttarting server on 4000")

	err := http.ListenAndServe(":4000", mux)

	log.Fatal(err)
}
