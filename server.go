// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 19.
//!+

// Server1 is a minimal "echo" server.
package main

import (
	//	"strings"
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
	"image/jpeg"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	server := &http.Server{
		Addr:           ":12345",
		Handler:        http.HandlerFunc(handler),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println("Listening on port 12345")
	log.Fatal(server.ListenAndServe())
}

var ImageTemplate string = `<!DOCTYPE html>
    <html lang="en"><head></head>
    <body><h1>Living Room Live Stream</h1><img width="60%" src="data:image/jpg;base64,{{.Image}}"></body>`

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	timeNow := time.Now().Format("Mon Jan _2 2006 15:04:05")

	// // Take a picture on the raspberry pi
	// cmd := exec.Command("raspistill", "-t", "1", "-q", "10", "-o", "www/img/most_recent.jpg")
	// err := cmd.Run()
	// if err != nil {
	// 	fmt.Println("command raspistill failed")
	// }
	//fmt.Println("Execed command at ", timeNow)
	//http.ServeFile(w, r, "./www/index.html")
	//http.FileServer(http.Dir("./www/")),

	imageFile, err := os.Open("www/img/most_recent.jpg")
	if err != nil {
		fmt.Println("Image reading error!")
	}
	defer imageFile.Close()

	loadedImage, err := jpeg.Decode(imageFile)
	if err != nil {
		fmt.Println("Decoding failed")
	}

	// In-memory buffer to store JPEG image
	// before we base 64 encode it
	var buff bytes.Buffer

	// The Buffer satisfies the Writer interface so we can use it with Encode
	// In previous example we encoded to a file, this time to a temp buffer
	jpeg.Encode(&buff, loadedImage, nil)

	// Encode the bytes in the buffer to a base64 string
	encodedString := base64.StdEncoding.EncodeToString(buff.Bytes())

	// You can embed it in an html doc with this string
	//htmlImage := "<img src=\"data:image/png;base64," + encodedString + "\" />"
	_ = timeNow
	_ = encodedString

	//w.Header().Set("Content-Length", strconv.Itoa(len(buff.Bytes())))
	//w.Header().Set("Cache-Control", "no-cache")
	//w.Header().Set("content-type", "text/html")

	tmpl, err := template.New("image").Parse(ImageTemplate)
	if err != nil {
		log.Println("unable to parse image template.")
	}
	data := map[string]interface{}{"Image": encodedString}
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println("unable to execute template.")
	}
}

//!-
