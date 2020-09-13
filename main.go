package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, "env:", os.Getenv("ECHO"))
		_, _ = fmt.Fprintln(w, "requestURI:", r.RequestURI)
		body, _ := ioutil.ReadAll(r.Body)
		_, _ = fmt.Fprintln(w, "body:", string(body))
		_, _ = fmt.Fprintln(w, "header:")
		for k, v := range r.Header {
			_, _ = fmt.Fprintln(w, "   ", k, ":", v)
		}
	})
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
