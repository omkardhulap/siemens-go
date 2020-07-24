package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// ask user number
	url := "https://reqres.in/api/users/" + "1"
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("Error handling Code for Get")
	} else {
		fmt.Println("in else", resp)
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error handling Code for ReadAll Resp body")
		} else {
			fmt.Println("body ", string(body))
		}
	}

}
