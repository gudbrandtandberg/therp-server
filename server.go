// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 19.
//!+

// Server1 is a minimal "echo" server.
package main

import (
	//	"strings"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	//	http.HandleFunc("/", handler) // each request calls handler
	//	http.Handle("/", http.FileServer(http.Dir("./www")))

	//	fmt.Println("HTTP server listening on port 12345")
	//	go func() {
	//		log.Fatal(http.ListenAndServe(":12345", nil))
	//	}()

	s := &http.Server{
		Addr:    ":12345",
		Handler: http.HandlerFunc(handler),
		//Handler:        http.FileServer(http.Dir("./www/")),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println("Listening on port 12345")
	log.Fatal(s.ListenAndServe())
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	timeNow := time.Now().Format("Mon Jan _2 2006 15:04:05")

	// // Take a picture on the raspberry pi
	// cmd := exec.Command("raspistill", "-t", "1", "-q", "10", "-o", "www/img/most_recent.jpg")
	// err := cmd.Run()
	// if err != nil {
	// 	fmt.Println("command raspistill failed")
	// }

	fmt.Println("Execed command at ", timeNow)
	http.ServeFile(w, r, "./www/index.html")
	//http.FileServer(http.Dir("./www/")),
}

//!-
