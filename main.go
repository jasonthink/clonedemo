package main

import (
	"fmt"
	"net/http"
)

func rootHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "<html><head><title>Test Clone and Reinstate</title><body><h1>Demo</h1><p><script type=\"text/javascript\">document.write(location.href);</script></p></body></html>")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/hello", rootHandler)
	http.HandleFunc("/headers", headers)
	http.ListenAndServe(":80", nil)
}
