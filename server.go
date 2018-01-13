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
	"time"
)

func main() {
	//	http.HandleFunc("/", handler) // each request calls handler
	http.Handle("/", http.FileServer(http.Dir("./www")))

	fmt.Println("HTTP server listening on port 12345")
	go func() {
		log.Fatal(http.ListenAndServe(":12345", nil))
	}()

	s := &http.Server{
		Addr:           ":12346",
		Handler:        http.HandlerFunc(handler),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println("Listening on port 12346")
	log.Fatal(s.ListenAndServe())
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, time.Now().Format("Mon Jan _2 2006 15:04:05"))
}

//!-
