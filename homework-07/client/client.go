package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			return
		}

		if r.Body != nil {
			defer r.Body.Close()
		}
		time.Sleep(time.Second * 1)

		body, readErr := ioutil.ReadAll(r.Body)
		if readErr != nil {
			log.Fatal(readErr)
		}

		var a string
		_ = json.Unmarshal(body, &a)
		fmt.Println(a)
		data, err := json.Marshal("pong")
		if err != nil {
			log.Fatal(err)
		}
		resp, err := http.Post("http://localhost:4000/", "application/json", bytes.NewReader(data))
		if err != nil {
			log.Fatal(err)
		}
		if resp != nil || resp.StatusCode != 200 {
			log.Fatal(err)
		}

		return

	})

	fmt.Println("Server has been started on port 4040")
	err := http.ListenAndServe(":4040", nil)
	if err != nil {
		log.Fatal(err)
	}
}
