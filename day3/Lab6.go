package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func main() {
	// ask user number
	var userno int
	fmt.Println("Enter user number : ")
	fmt.Scanf("%d", &userno)
	url := "https://reqres.in/api/users/" + strconv.Itoa(userno)
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("Error handling Code for Get")
	} else {
		body, _ := ioutil.ReadAll(resp.Body)

		fmt.Println("body ", string(body))

	}

}
