// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 19.
//!+

// Server1 is a minimal "echo" server.
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
//	http.HandleFunc("/", handler) // each request calls handler
	http.Handle("/", http.FileServer(http.Dir("./www")))
	fmt.Println("HTTP server listening on port 12345")
	log.Fatal(http.ListenAndServe(":12345", nil))
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	fmt.Fprintf(w, "This is our site")
}

//!-

