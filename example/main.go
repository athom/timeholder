package main

import (
	"github.com/athom/timeholder"
	"io"
	"log"
	"net/http"
	"time"
)

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	key := req.FormValue("key")
	timeholder.SetDuration(5 * time.Second)
	if timeholder.Hold(key) {
		io.WriteString(w, "Be patient!")
		return
	}

	io.WriteString(w, "hello, world!\n")
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
