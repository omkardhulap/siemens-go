package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func process() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "<h1>Main Page</h1>")
	}
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/emp", Emphandler)
	fmt.Printf("sever starting on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func main() {
	process()
}
