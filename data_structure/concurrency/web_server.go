package main

import (
	"io"
	"log"
	"net/http"
)

const form = `<html><body><form action="#" method="post" name="bar">
<input type="text" name="in"/>
<input type="submit" value="Submit"/>
</form></body></html>`

func SimpleServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "<h1>Hello World!</h1>")
}

func FormServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	switch req.Method {
	case "GET":
		io.WriteString(w, form)
	case "POST":
		io.WriteString(w, req.FormValue("in"))
	}
}

func main() {
	http.HandleFunc("/test1", SimpleServer)
	http.HandleFunc("/test2", FormServer)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
