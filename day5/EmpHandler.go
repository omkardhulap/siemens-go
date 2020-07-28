package main

import (
	"io"
	"net/http"
)

func Emphandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		io.WriteString(w, "<h1>GET</h1>")
	case "POST":
	  	  io.WriteString(w, "<h1>EmpHandler POST Page</h1>")
	    }
}
