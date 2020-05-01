package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var AuthBaseURL = os.Getenv("AUTH_BASE_URL")

func main() {
	http.HandleFunc("/post", post)

	http.ListenAndServe(":8092", nil)
}

func post(w http.ResponseWriter, req *http.Request) {
	resp, err := http.Get(AuthBaseURL + "/auth")
	if err != nil {
		log.Println("failed to get auth ", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	fmt.Fprintf(w, "Post : "+string(body))
}
