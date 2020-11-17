package main

import (
	"fmt"
	"log"
	"net/http"
)

func rootHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("hello called")
	fmt.Fprintf(w, "<html><head><title>Test Clone and Reinstate</title><body><h1>Demo</h1><p><script type=\"text/javascript\">document.write(location.href);</script></p></body></html>")
}

func headers(w http.ResponseWriter, req *http.Request) {
	log.Println("/headers requested")
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	log.Println("Starting")
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/hello", rootHandler)
	http.HandleFunc("/headers", headers)
	log.Println("Listening on 3000")
	http.ListenAndServe(":3000", nil)
	log.Println("Exiting")
}
