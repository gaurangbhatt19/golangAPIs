package main

import (
	"fmt"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "<h1>HomePage</h1>")
}

func main() {
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":3033", nil)
}
