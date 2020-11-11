package main

import (
    "fmt"
		"net/http"
		"os"
)

func hostname(w http.ResponseWriter, req *http.Request) {
	hostname, err := os.Hostname()
    if err != nil {
        panic(err)
    }
    fmt.Fprintf(w, "Serving from %s\n", hostname)
}

func post(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "post request accepted\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func main() {
    http.HandleFunc("/", hostname)
    http.HandleFunc("/data", post)
		http.HandleFunc("/headers", headers)
		
		fmt.Println("Starting to serve from :8080 ...")

    http.ListenAndServe(":8080", nil)
}
