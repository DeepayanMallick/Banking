package main

import (
	"fmt"
	"net/http"
)

func main()  {
	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world")
	})

	http.ListenAndServe("localhost:8000", nil)
}