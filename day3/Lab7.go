package main

import (
	"fmt"
	"io"
	http "net/http"
)

func main() {
	myHttpHandler := func(
		respwriter http.ResponseWriter,
		req *http.Request) {
		io.WriteString(respwriter,
			"<H1>In HTTPHandler of Go</h1>")
	}
	myempHandler := func(
		respwriter http.ResponseWriter,
		req *http.Request) {
		io.WriteString(respwriter,
			"<H1>In HTTPHandler of /emp</h1>")
	}

	http.HandleFunc("/", myHttpHandler)
	http.HandleFunc("/emp", myempHandler)
	fmt.Println("Starting server on 8085")
	err := http.ListenAndServe("localhost:8085", nil)

	fmt.Println("err", err)
	/*
		http.HandleFunc("/", helloHandler)
		fmt.Printf("sever starting on 8080")
		log.Fatal(http.ListenAndServe(":8080", nil))
	*/
}
