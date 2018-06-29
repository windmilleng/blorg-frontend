package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"flag"
	"io/ioutil"
	"strings"
)

const endptPong = "/pong"

var backend = flag.String("backendAddr", "http://localhost:8080", "address of the blorg backend server")

func main() {
	flag.Parse()

	r := mux.NewRouter()
	r.HandleFunc("/", Hello)
	r.HandleFunc("/ping", Ping)
	http.Handle("/", r)
	fmt.Println("Starting up on 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "I'm the frontend server.")
}

func Ping(w http.ResponseWriter, req *http.Request) {
	// TODO(maia): Will want to be more careful concat'ing base + endpt in future
	// see http://bit.ly/2lFlOCq
	url := fmt.Sprintf("%s%s", *backend, endptPong)
	resp, err := http.Get(url)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Request to backend server resulted in error: %v\n", err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Could not read backend server resp. with error: %v\n", err)
		return
	}

	if resp.StatusCode != 200 || strings.TrimSpace(string(body)) != "pong" {
		w.WriteHeader(500)
		fmt.Fprintf(w,
			"Expected 'pong'; backend server responded with 'pong' (status: %s)\n", resp.Status)
		return
	}

	fmt.Fprintf(w, "SUCCESS! ﾍ(=￣∇￣)ﾉ\n")
}
