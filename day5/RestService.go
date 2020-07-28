package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func loadPage(title string) (*[]byte, error) {
	fmt.Println(title)
	body, err := ioutil.ReadFile(title)
	if err != nil {
		return nil, err
	}
	return &body, nil
}

func process() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		body, _ := loadPage("post.html")
		fmt.Fprintln(w, string(*body))
	}
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/emp", Emphandler)
	fmt.Printf("sever starting on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func main() {
	process()
}
