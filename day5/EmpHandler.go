package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type EmpStruct struct {
	Empno  int     `json:"empno"`
	Ename  string  `json:"ename"`
	Salary float64 `json:"salary"`
}

func Emphandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		io.WriteString(w, "<h1>GET</h1>")
	case "POST":
		fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", req.PostForm)
		reqBody, _ := ioutil.ReadAll(req.Body)
		var emp EmpStruct
		json.Unmarshal(reqBody, &emp)
		Save(emp)
		json.NewEncoder(w).Encode(emp)
	}
}
