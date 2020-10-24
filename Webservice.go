package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {

	response := make(chan *http.Response, 1)
	errors := make(chan *error)
	go func() {
		res, err := http.Get("https://jsonplaceholder.typicode.com/posts/42")
		if err != nil {
			errors <- &err
		}
		response <- res

	}()

	for {
		select {
		case r := <-response:
			fmt.Println("Response Received")
			defer r.Body.Close()
			body, _ := ioutil.ReadAll(r.Body)
			fmt.Println("_________")
			fmt.Println(string(body))
			//fmt.Printf("%s", body)

			return
		case err := <-errors:
			fmt.Println("Error Received")
			log.Fatal(*err)
			return

		case <-time.After(2000 * time.Microsecond):
			fmt.Println("Timeout Error")
			return
		}
	}
}
