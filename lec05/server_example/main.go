package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprint(w, "Hello from go web server", r.URL.Path[1:])
	})
	http.ListenAndServe(":8080", nil)
}
