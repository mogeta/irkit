package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mogeta/irkit/relay"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

type Input struct {
	In string
}

type Output struct {
	Out string
}

func jsonHandleFunc(rw http.ResponseWriter, req *http.Request) {

	output := Output{relay.GetIPAddress()}
	outjson, err := json.Marshal(output)
	if err != nil {
		fmt.Println(err)
	}
	rw.Header().Set("Content-Type", "application/json")
	fmt.Fprint(rw, string(outjson))
}

func main() {
	relay.GetIPAddress()

	http.HandleFunc("/", handler)
	http.HandleFunc("/api", jsonHandleFunc)
	http.ListenAndServe(":8080", nil)
}
