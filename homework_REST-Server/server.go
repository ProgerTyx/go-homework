package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = ioutil.ReadAll(r.Body)
		fmt.Println(r.Method, r.Body)
		fmt.Printf("Hello")
	})
	_ = http.ListenAndServe(":4000", nil)
}
