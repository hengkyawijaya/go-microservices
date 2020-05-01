package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/auth", auth)

	http.ListenAndServe(":8091", nil)
}

func auth(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "authenticated")
}
