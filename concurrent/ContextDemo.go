package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		context := request.Context()
		fmt.Println("Hello server start!")
		defer fmt.Println("hello server ended!")
		select {
		case <-time.After(10 * time.Second):
			fmt.Fprintf(writer, "hello/n")
		case <-context.Done():
			err := context.Err()
			fmt.Println("hello server:", err)
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	})
	http.ListenAndServe(":8090",nil)

}
