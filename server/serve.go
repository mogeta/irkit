package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mogeta/irkit/relay"
)

var irkit *relay.Irkit

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

type Input struct {
	In string
}

type Output struct {
	Out string
}

type Format struct {
	Ok           int
	ErrorMessage string
}

type SignalData struct {
	Format
	Data string
}

func jsonHandleFunc(rw http.ResponseWriter, req *http.Request) {

	output := Output{irkit.GetIPAddress()}
	outjson, err := json.Marshal(output)
	if err != nil {
		fmt.Println(err)
	}
	rw.Header().Set("Content-Type", "application/json")
	fmt.Fprint(rw, string(outjson))
}

func message(rw http.ResponseWriter, req *http.Request) {
	signal, err := irkit.GetMessages()
	result := Format{0, ""}
	if err != nil {
		result.Ok = 1
		result.ErrorMessage = "irkitError"
	}
	output := SignalData{result, signal}
	outjson, err := json.Marshal(output)
	if err != nil {
		fmt.Println(err)
	}
	rw.Header().Set("Content-Type", "application/json")
	fmt.Fprint(rw, string(outjson))
}

func main() {
	irkit = relay.New()

	http.HandleFunc("/", handler)
	http.HandleFunc("/api", jsonHandleFunc)
	http.HandleFunc("/message", message)
	http.ListenAndServe(":8080", nil)
}
